//document url:https://www.ibm.com/developerworks/cloud/library/cl-ibm-blockchain-chaincode-development-using-golang/index.html


package main
import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//custom data models
type PersonalInfo struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    DOB       string `json:"DOB"`
    Email     string `json:"email"`
    Mobile    string `json:"mobile"`
}
 
type FinancialInfo struct {
    MonthlySalary      int `json:"monthlySalary"`
    MonthlyRent        int `json:"monthlyRent"`
    OtherExpenditure   int `json:"otherExpenditure"`
    MonthlyLoanPayment int `json:"monthlyLoanPayment"`
}
 
type LoanApplication struct {
    ID                     string        `json:"id"`
    PropertyID             string        `json:"propertyId"`
    LandID                 string        `json:"landId"`
    PermitID               string        `json:"permitId"`
    BuyerID                string        `json:"buyerId"`
    SalesContractID        string        `json:"salesContractId"`
    PersonalInfo           PersonalInfo  `json:"personalInfo"`
    FinancialInfo          FinancialInfo `json:"financialInfo"`
    Status                 string        `json:"status"`
    RequestedAmount        int           `json:"requestedAmount"`
    FairMarketValue        int           `json:"fairMarketValue"`
    ApprovedAmount         int           `json:"approvedAmount"`
    ReviewerID             string        `json:"reviewerId"`
    LastModifiedDate       string        `json:"lastModifiedDate"`
}


type SampleChaincode struct {
}
 
func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}


func main(){
err := shim.Start(new(SampleChaincode))
    if err != nil {
        fmt.Println("Could not start SampleChaincode")
    } else {
        fmt.Println("SampleChaincode successfully started")
    }
 
}