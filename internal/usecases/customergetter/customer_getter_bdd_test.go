package customergetter_test

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	gomock "go.uber.org/mock/gomock"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	mock_repository "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/customerrepositorymock"
	customergetter "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customergetter"
)

var (
	ctx              context.Context
	mockCtrl         *gomock.Controller
	getterUseCase    *customergetter.CustomerGetter
	mockCustomerRepo *mock_repository.MockCustomerRepository
	customer         *entities.Customer
	err              error
)

func TestMain(t *testing.T) {
	ctx = context.TODO()
	mockCtrl = gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCustomerRepo = mock_repository.NewMockCustomerRepository(mockCtrl)
	getterUseCase = customergetter.NewCustomerGetter(mockCustomerRepo)

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the customer with document "([^"]*)" exists$`, theCustomerWithDocumentExists)
	ctx.Step(`^I retrieve the customer details$`, iRetrieveTheCustomerDetails)
	ctx.Step(`^I should get the customer details$`, iShouldGetTheCustomerDetails)
}

func theCustomerWithDocumentExists(document string) error {
	customer = &entities.Customer{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Document: document,
	}
	mockCustomerRepo.EXPECT().GetByDocument(ctx, document).Return(customer, nil)
	return nil
}

func iRetrieveTheCustomerDetails() error {
	_, err = getterUseCase.Get(ctx, customer)
	return err
}

func iShouldGetTheCustomerDetails() error {
	if err != nil {
		return err
	}
	return nil
}
