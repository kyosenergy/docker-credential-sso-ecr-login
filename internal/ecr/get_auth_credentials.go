package ecr

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ssocreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/sts"

	log "github.com/sirupsen/logrus"
)

var (
	HelperName = "sso-ecr-login"

	registryURL string
	accountID   string
	region      string
	username    string
	secret      string
	sess        *session.Session
)

// Parses the account ID and region from the registry URL when docker calls this helper.
func getAccountAndRegionFromRequest() error {
	_, err := fmt.Scanln(&registryURL)
	if err != nil {
		return fmt.Errorf("error parsing registry URL: %s", err)
	}
	log.Infof("[%s] Registry URL: %s", HelperName, registryURL)

	// Parse account ID and region from the server URL
	parts := strings.Split(registryURL, ".")
	if len(parts) < 4 {
		return fmt.Errorf("error parsing registry URL: %s", registryURL)
	}

	accountID = parts[0]
	log.Debugf("[%s] Account ID: %s", HelperName, accountID)
	region = parts[3]
	log.Debugf("[%s] Region: %s", HelperName, region)

	return nil
}

// Check if AWS SSO login is needed
func ensureNoSSOLoginNeeded() error {
	profile := os.Getenv("AWS_PROFILE")
	if profile == "" {
		profile = "default"

		log.Infof("[%s] AWS_PROFILE is not set. Will use a [default] profile", HelperName)
	}

	sess, _ = session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	svc := sts.New(sess)
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)
	if errors.Is(err, credentials.ErrNoValidProvidersFoundInChain) {
		return fmt.Errorf("no configuration found for profile: %s", profile)
	}

	if err != nil {
		var awsError awserr.Error
		if errors.As(err, &awsError) {
			switch awsError.Code() {
			case ssocreds.ErrCodeSSOProviderInvalidToken:
				return fmt.Errorf("invalid SSO token. Please run 'aws sso login' to refresh the token")
			default:
				return fmt.Errorf("error: %s", awsError.Error())
			}
		}
	}

	if *result.Account != accountID {
		return fmt.Errorf("the account ID %s does not match the expected account ID %s", *result.Account, accountID)
	}

	log.Infof("[%s] Using AWS SSO session: %s", HelperName, *result.Arn)

	return nil
}

// GetAuthCredentials the ECR login password using the AWS SDK
func getECRLoginPassword() error {
	ecrClient := ecr.New(sess)

	// GetAuthCredentials login password
	resp, err := ecrClient.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return fmt.Errorf("failed to get authorization token: %v", err)
	}

	// Assume the first authorization token is what we need
	if len(resp.AuthorizationData) == 0 {
		return fmt.Errorf("no authorization data found")
	}

	// Extract the password from the authorization token
	authorizationData := resp.AuthorizationData[0]
	authToken := authorizationData.AuthorizationToken

	// Decode the base64 encoded token
	decodedToken, err := base64.StdEncoding.DecodeString(*authToken)
	if err != nil {
		return fmt.Errorf("failed to decode authorization token: %v", err)
	}

	// The token is in the form "username:secret"
	parts := strings.SplitN(string(decodedToken), ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid authorization token format")
	}

	username, secret = parts[0], parts[1]

	return nil
}

// Output the credentials in Docker JSON format
func outputDockerJsonLoginFormat() error {
	// Output in Docker JSON format
	dockerCredentials := map[string]string{
		"ServerURL": registryURL,
		"Username":  username,
		"Secret":    secret,
	}

	responseForDocker, err := json.Marshal(dockerCredentials)
	if err != nil {
		return fmt.Errorf("error marshalling docker credentials: %s", err)
	}

	fmt.Println(string(responseForDocker))
	return nil
}

func GetAuthCredentials() {
	handleError(getAccountAndRegionFromRequest())
	handleError(ensureNoSSOLoginNeeded())
	handleError(getECRLoginPassword())
	handleError(outputDockerJsonLoginFormat())
}
