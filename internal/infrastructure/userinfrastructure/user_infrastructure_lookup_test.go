package userinfrastructure

import (
	"context"
	"os"
	"testing"

	"github.com/Hirochon/infra-go-for-test/internal/domain/user"
	"github.com/Hirochon/infra-go-for-test/internal/infrastructure/externalconnection/mysqlconnection"
	"github.com/Hirochon/infra-go-for-test/internal/infrastructure/sqlc/sqlcgentestmodel"
)

type ScenarioKey string

func testUserRepositoryLookUpSuccess(ctx context.Context, t *testing.T, userRepository user.IUserRepository) {
	cases := []struct {
		scenario string
		id       string
	}{
		{
			scenario: "正常なユーザーIDを渡した場合、ユーザーIDが返却される①",
			id:       "01H0YHVK6NREDFDEK81YT1ERKR",
		},
		{
			scenario: "正常なユーザーIDを渡した場合、ユーザーIDが返却される②",
			id:       "01H0YHVK6MK92QZ1SZD2HAXEZ2",
		},
		{
			scenario: "正常なユーザーIDを渡した場合、ユーザーIDが返却される③",
			id:       "01H0YHVK6N4DJVRX7HXS9KSF4V",
		},
	}
	for _, c := range cases {
		cc := c
		t.Run(cc.scenario, func(t *testing.T) {
			t.Parallel()
			ctx := context.WithValue(ctx, ScenarioKey("testScenario"), cc.scenario)
			_, err := userRepository.LookUp(ctx, cc.id)
			if err != nil {
				t.Errorf("エラーが発生した: %v", err)
			}
		})
	}
}

func testUserRepositoryLookUpFailed(ctx context.Context, t *testing.T, userRepository user.IUserRepository) {
	cases := []struct {
		scenario string
		id       string
	}{
		{
			scenario: "異常なユーザーIDを渡した場合、エラーが発生する①",
			id:       "01H0YHVK6NREDFDEK81YT1ERKR1",
		},
		{
			scenario: "異常なユーザーIDを渡した場合、エラーが発生する②",
			id:       "01H0YHVK6MK92QZ1SZD2HAXEZ",
		},
		{
			scenario: "異常なユーザーIDを渡した場合、エラーが発生する③",
			id:       "",
		},
	}
	for _, c := range cases {
		cc := c
		t.Run(cc.scenario, func(t *testing.T) {
			t.Parallel()
			ctx := context.WithValue(ctx, ScenarioKey("testScenario"), cc.scenario)
			_, err := userRepository.LookUp(ctx, cc.id)
			if err == nil {
				t.Errorf("エラーが発生しなかった: %v", err)
			}
		})
	}
}

func TestUserRepositoryLookUp(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mysqlClient, err := mysqlconnection.NewMySQLClient(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_EXTRA_PROPERTIES"))
	if err != nil {
		t.Fatalf("予期しないエラー: %v", err)
	}
	userRepository := NewUserRepository(mysqlClient)

	testDataGenerator := sqlcgentestmodel.New(mysqlClient)
	testDataGenerator.TestCreateUser(ctx)
	t.Cleanup(func() {
		testDataGenerator.TestDeleteUser(ctx)
	})

	testUserRepositoryLookUpSuccess(ctx, t, userRepository)
	testUserRepositoryLookUpFailed(ctx, t, userRepository)
}
