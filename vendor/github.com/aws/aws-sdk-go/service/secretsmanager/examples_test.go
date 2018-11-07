// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To cancel scheduled rotation for a secret
//
// The following example shows how to cancel rotation for a secret. The operation sets
// the RotationEnabled field to false and cancels all scheduled rotations. To resume
// scheduled rotations, you must re-enable rotation by calling the rotate-secret operation.
func ExampleSecretsManager_CancelRotateSecret_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.CancelRotateSecretInput{
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.CancelRotateSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a basic secret
//
// The following example shows how to create a secret. The credentials stored in the
// encrypted secret value are retrieved from a file on disk named mycreds.json.
func ExampleSecretsManager_CreateSecret_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.CreateSecretInput{
		ClientRequestToken: aws.String("EXAMPLE1-90ab-cdef-fedc-ba987SECRET1"),
		Description:        aws.String("My test database secret created with the CLI"),
		Name:               aws.String("MyTestDatabaseSecret"),
		SecretString:       aws.String("{\"username\":\"david\",\"password\":\"BnQw!XDWgaEeT9XGTT29\"}"),
	}

	result, err := svc.CreateSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeEncryptionFailure:
				fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeResourceExistsException:
				fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodePreconditionNotMetException:
				fmt.Println(secretsmanager.ErrCodePreconditionNotMetException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete the resource-based policy attached to a secret
//
// The following example shows how to delete the resource-based policy that is attached
// to a secret.
func ExampleSecretsManager_DeleteResourcePolicy_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.DeleteResourcePolicyInput{
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.DeleteResourcePolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a secret
//
// The following example shows how to delete a secret. The secret stays in your account
// in a deprecated and inaccessible state until the recovery window ends. After the
// date and time in the DeletionDate response field has passed, you can no longer recover
// this secret with restore-secret.
func ExampleSecretsManager_DeleteSecret_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.DeleteSecretInput{
		RecoveryWindowInDays: aws.Int64(7),
		SecretId:             aws.String("MyTestDatabaseSecret1"),
	}

	result, err := svc.DeleteSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve the details of a secret
//
// The following example shows how to get the details about a secret.
func ExampleSecretsManager_DescribeSecret_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.DescribeSecretInput{
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.DescribeSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To generate a random password
//
// The following example shows how to request a randomly generated password. This example
// includes the optional flags to require spaces and at least one character of each
// included type. It specifies a length of 20 characters.
func ExampleSecretsManager_GetRandomPassword_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetRandomPasswordInput{
		IncludeSpace:            aws.Bool(true),
		PasswordLength:          aws.Int64(20),
		RequireEachIncludedType: aws.Bool(true),
	}

	result, err := svc.GetRandomPassword(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve the resource-based policy attached to a secret
//
// The following example shows how to retrieve the resource-based policy that is attached
// to a secret.
func ExampleSecretsManager_GetResourcePolicy_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetResourcePolicyInput{
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.GetResourcePolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve the encrypted secret value of a secret
//
// The following example shows how to retrieve the secret string value from the version
// of the secret that has the AWSPREVIOUS staging label attached. If you want to retrieve
// the AWSCURRENT version of the secret, then you can omit the VersionStage parameter
// because it defaults to AWSCURRENT.
func ExampleSecretsManager_GetSecretValue_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String("MyTestDatabaseSecret"),
		VersionStage: aws.String("AWSPREVIOUS"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeDecryptionFailure:
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list all of the secret versions associated with a secret
//
// The following example shows how to retrieve a list of all of the versions of a secret,
// including those without any staging labels.
func ExampleSecretsManager_ListSecretVersionIds_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.ListSecretVersionIdsInput{
		IncludeDeprecated: aws.Bool(true),
		SecretId:          aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.ListSecretVersionIds(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidNextTokenException:
				fmt.Println(secretsmanager.ErrCodeInvalidNextTokenException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list the secrets in your account
//
// The following example shows how to list all of the secrets in your account.
func ExampleSecretsManager_ListSecrets_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.ListSecretsInput{}

	result, err := svc.ListSecrets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidNextTokenException:
				fmt.Println(secretsmanager.ErrCodeInvalidNextTokenException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add a resource-based policy to a secret
//
// The following example shows how to add a resource-based policy to a secret.
func ExampleSecretsManager_PutResourcePolicy_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.PutResourcePolicyInput{
		ResourcePolicy: aws.String("{\n\"Version\":\"2012-10-17\",\n\"Statement\":[{\n\"Effect\":\"Allow\",\n\"Principal\":{\n\"AWS\":\"arn:aws:iam::123456789012:root\"\n},\n\"Action\":\"secretsmanager:GetSecretValue\",\n\"Resource\":\"*\"\n}]\n}"),
		SecretId:       aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.PutResourcePolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To store a secret value in a new version of a secret
//
// The following example shows how to create a new version of the secret. Alternatively,
// you can use the update-secret command.
func ExampleSecretsManager_PutSecretValue_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.PutSecretValueInput{
		ClientRequestToken: aws.String("EXAMPLE2-90ab-cdef-fedc-ba987EXAMPLE"),
		SecretId:           aws.String("MyTestDatabaseSecret"),
		SecretString:       aws.String("{\"username\":\"david\",\"password\":\"BnQw!XDWgaEeT9XGTT29\"}"),
	}

	result, err := svc.PutSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeEncryptionFailure:
				fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeResourceExistsException:
				fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To restore a previously deleted secret
//
// The following example shows how to restore a secret that you previously scheduled
// for deletion.
func ExampleSecretsManager_RestoreSecret_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.RestoreSecretInput{
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.RestoreSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To configure rotation for a secret
//
// The following example configures rotation for a secret by providing the ARN of a
// Lambda rotation function (which must already exist) and the number of days between
// rotation. The first rotation happens immediately upon completion of this command.
// The rotation function runs asynchronously in the background.
func ExampleSecretsManager_RotateSecret_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.RotateSecretInput{
		RotationLambdaARN: aws.String("arn:aws:lambda:us-west-2:123456789012:function:MyTestDatabaseRotationLambda"),
		RotationRules: &secretsmanager.RotationRulesType{
			AutomaticallyAfterDays: aws.Int64(30),
		},
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.RotateSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To request an immediate rotation for a secret
//
// The following example requests an immediate invocation of the secret's Lambda rotation
// function. It assumes that the specified secret already has rotation configured. The
// rotation function runs asynchronously in the background.
func ExampleSecretsManager_RotateSecret_shared01() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.RotateSecretInput{
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.RotateSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add tags to a secret
//
// The following example shows how to attach two tags each with a Key and Value to a
// secret. There is no output from this API. To see the result, use the DescribeSecret
// operation.
func ExampleSecretsManager_TagResource_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.TagResourceInput{
		SecretId: aws.String("MyExampleSecret"),
		Tags: []*secretsmanager.Tag{
			{
				Key:   aws.String("FirstTag"),
				Value: aws.String("SomeValue"),
			},
			{
				Key:   aws.String("SecondTag"),
				Value: aws.String("AnotherValue"),
			},
		},
	}

	result, err := svc.TagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To remove tags from a secret
//
// The following example shows how to remove two tags from a secret's metadata. For
// each, both the tag and the associated value are removed. There is no output from
// this API. To see the result, use the DescribeSecret operation.
func ExampleSecretsManager_UntagResource_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.UntagResourceInput{
		SecretId: aws.String("MyTestDatabaseSecret"),
		TagKeys: []*string{
			aws.String("FirstTag"),
			aws.String("SecondTag"),
		},
	}

	result, err := svc.UntagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update the description of a secret
//
// The following example shows how to modify the description of a secret.
func ExampleSecretsManager_UpdateSecret_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.UpdateSecretInput{
		ClientRequestToken: aws.String("EXAMPLE1-90ab-cdef-fedc-ba987EXAMPLE"),
		Description:        aws.String("This is a new description for the secret."),
		SecretId:           aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.UpdateSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeEncryptionFailure:
				fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeResourceExistsException:
				fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodePreconditionNotMetException:
				fmt.Println(secretsmanager.ErrCodePreconditionNotMetException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update the KMS key associated with a secret
//
// This example shows how to update the KMS customer managed key (CMK) used to encrypt
// the secret value. The KMS CMK must be in the same region as the secret.
func ExampleSecretsManager_UpdateSecret_shared01() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.UpdateSecretInput{
		KmsKeyId: aws.String("arn:aws:kms:us-west-2:123456789012:key/EXAMPLE2-90ab-cdef-fedc-ba987EXAMPLE"),
		SecretId: aws.String("MyTestDatabaseSecret"),
	}

	result, err := svc.UpdateSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeEncryptionFailure:
				fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeResourceExistsException:
				fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodePreconditionNotMetException:
				fmt.Println(secretsmanager.ErrCodePreconditionNotMetException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new version of the encrypted secret value
//
// The following example shows how to create a new version of the secret by updating
// the SecretString field. Alternatively, you can use the put-secret-value operation.
func ExampleSecretsManager_UpdateSecret_shared02() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String("MyTestDatabaseSecret"),
		SecretString: aws.String("{JSON STRING WITH CREDENTIALS}"),
	}

	result, err := svc.UpdateSecret(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeEncryptionFailure:
				fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeResourceExistsException:
				fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodePreconditionNotMetException:
				fmt.Println(secretsmanager.ErrCodePreconditionNotMetException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add a staging label attached to a version of a secret
//
// The following example shows you how to add a staging label to a version of a secret.
// You can review the results by running the operation ListSecretVersionIds and viewing
// the VersionStages response field for the affected version.
func ExampleSecretsManager_UpdateSecretVersionStage_shared00() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.UpdateSecretVersionStageInput{
		MoveToVersionId: aws.String("EXAMPLE1-90ab-cdef-fedc-ba987SECRET1"),
		SecretId:        aws.String("MyTestDatabaseSecret"),
		VersionStage:    aws.String("STAGINGLABEL1"),
	}

	result, err := svc.UpdateSecretVersionStage(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a staging label attached to a version of a secret
//
// The following example shows you how to delete a staging label that is attached to
// a version of a secret. You can review the results by running the operation ListSecretVersionIds
// and viewing the VersionStages response field for the affected version.
func ExampleSecretsManager_UpdateSecretVersionStage_shared01() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.UpdateSecretVersionStageInput{
		RemoveFromVersionId: aws.String("EXAMPLE1-90ab-cdef-fedc-ba987SECRET1"),
		SecretId:            aws.String("MyTestDatabaseSecret"),
		VersionStage:        aws.String("STAGINGLABEL1"),
	}

	result, err := svc.UpdateSecretVersionStage(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To move a staging label from one version of a secret to another
//
// The following example shows you how to move a staging label that is attached to one
// version of a secret to a different version. You can review the results by running
// the operation ListSecretVersionIds and viewing the VersionStages response field for
// the affected version.
func ExampleSecretsManager_UpdateSecretVersionStage_shared02() {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.UpdateSecretVersionStageInput{
		MoveToVersionId:     aws.String("EXAMPLE2-90ab-cdef-fedc-ba987SECRET2"),
		RemoveFromVersionId: aws.String("EXAMPLE1-90ab-cdef-fedc-ba987SECRET1"),
		SecretId:            aws.String("MyTestDatabaseSecret"),
		VersionStage:        aws.String("AWSCURRENT"),
	}

	result, err := svc.UpdateSecretVersionStage(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
