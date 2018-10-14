package main


import (
    // "bytes"
     "encoding/json"
     "fmt"
     "math/rand"
     "strconv"

     "github.com/hyperledger/fabric/core/chaincode/shim"
     sc "github.com/hyperledger/fabric/protos/peer"

)

var cropInfo map[string]bool
 // Define the Smart Contract structure
type SmartContract struct {

}

 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library


type PackagedDetails struct {
     BatchId string `json:"batch"`                /* also known as Crop ID */
     BerrySize string `json:"berrysize"`
     NoOfCarton int `json:noofcarton`
     size int `json:size` 
 }

 type PalletDetails struct {
    PalletNo string `json:"palletno"`
    SSCCode string `json:"SSCCode"`
    PalletType string `json:"pallettype"`
    Quantity int `json:quantity`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
     return shim.Success(nil)
}

 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

     // Retrieve the requested Smart Contract function and arguments
     function, args := APIstub.GetFunctionAndParameters()
     // Route to the appropriate handler function to interact with the ledger appropriately
     if function == "viewPallotsReport" {
         return s.viewPallotsReport(APIstub, args)
     } else if function == "packageItemEntry" {
         return s.packageItemEntry(APIstub, args)
     } else if function == "uploadIpfsHash" {
            return s.uploadIpfsHash(APIstub, args)
     }

     return shim.Error("Invalid Smart Contract function name.")
}

 func (s *SmartContract) uploadIpfsHash(APIstub shim.ChaincodeStubInterface,args []string) sc.Response {
     if len(args) != 3 {
         return shim.Error("Incorrect number of arguments. Expecting IPFS Hash and Crop ID")
     }        

     var newReportInfo = args[1]
     cropInfo[newReportInfo] = true
     ipfsHashAsByte, _ := json.Marshal(newReportInfo)
     APIstub.PutState(args[0],ipfsHashAsByte)

     return shim.Success(nil)
}
 

func (s *SmartContract) packageItemEntry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
     if len(args) != 5 {
         return shim.Error("Incorrect number of arguments. Expecting 5-Those are CropID,BerrySize,NoOfCarton & size")
     }

    CropIdAsBytes, err := APIstub.GetState(args[1])
    if err != nil {
        return shim.Error("Failed at validation stage. The Crop ID does not exist/valid")
    }

    tempCropId := string(CropIdAsBytes)
    if exists := cropInfo[tempCropId]; !exists {
        return shim.Error("The Crop id does NOT exist!")
    }

    return s.createPallot(APIstub, args)              /* arguments CropID,noOfPallot,size*/
}

func (s *SmartContract) createPallot(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
     if len(args) != 5 {
         return shim.Error("Incorrect number of arguments. Expecting 3. Those are CropId,size, NoOfPallots")
     }
     //countOfPallots := args[4]/args[3]  

     countOfPallots := 5
    for i := 0; i <= countOfPallots; i++ {
     //for _, i := range countOfPallots {
        newPalNumber := rand.Intn(3520)
        newPallots := PalletDetails{PalletNo:strconv.Itoa(newPalNumber), SSCCode:"PallotOne", PalletType:"TypeOne", Quantity:20}
        pallotsAsBytes, _ := json.Marshal(newPallots)
        APIstub.PutState(args[0],pallotsAsBytes)
    }
    return shim.Success(nil)
}

func (s *SmartContract) viewPallotsReport(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    pallotsAsBytes, _ := APIstub.GetState(args[0])
    return shim.Success(pallotsAsBytes)
}

 // The main function is only relevant in unit test mode. Only included here for completeness.
func main() { 
     // Create a new Smart Contract
     err := shim.Start(new(SmartContract))
     if err != nil {
         fmt.Printf("Error creating new Smart Contract: %s", err)
     }
}
