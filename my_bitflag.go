package main

import "fmt"

func main() {
	for CreditAllocations := int32(0); CreditAllocations < 16; CreditAllocations++ {
		CreditAllocationStatus := ExtractCreditAllocationStatus(CreditAllocations)

		fmt.Printf("[%v]\n", CreditAllocations)
		for _, v := range CreditAllocationStatus {
			fmt.Printf("[%v]\n", v)
		}
		fmt.Printf("\n\n")
	}
}

func ExtractCreditAllocationStatus(CreditAllocations int32) []string {
	creditAllocationStatus := []string{}

	if extractBit(CreditAllocations, 0) == 1 {
		creditAllocationStatus = append(creditAllocationStatus, string(creditAllocationBitMapping[0]))
	}
	if extractBit(CreditAllocations, 1) == 1 {
		creditAllocationStatus = append(creditAllocationStatus, string(creditAllocationBitMapping[1]))
	}
	if extractBit(CreditAllocations, 2) == 1 {
		creditAllocationStatus = append(creditAllocationStatus, string(creditAllocationBitMapping[2]))
	}
	if extractBit(CreditAllocations, 3) == 1 {
		creditAllocationStatus = append(creditAllocationStatus, string(creditAllocationBitMapping[3]))
	}

	return creditAllocationStatus
}

func extractBit(num int32, position uint) int32 {
	return (num >> position) & 1 // Extract the bit (0 or 1) from the position.
}

var creditAllocationBitMapping = map[int32]creditAllocation{
	0: endpointSensorDetectionAndResponse,
	1: advancedEndpointSecurity,
	2: advancedServerAndWorkloadProtection,
	3: sapScannerForVisionOneEndpointSecurityPro,
}

type creditAllocation string

const ( // Should maintain wording with HIE figma: https://www.figma.com/design/Qc88PAh9wdodfzdUwV5F82/Server-%26-Workload-Protection?node-id=4871-18491&t=7ZS9JJIrCiPsoqHQ-0
	endpointSensorDetectionAndResponse        creditAllocation = "Endpoint sensor detection and response (1)" // [TODO]: remove brackets
	advancedEndpointSecurity                  creditAllocation = "Advanced Endpoint Security (2)"
	advancedServerAndWorkloadProtection       creditAllocation = "Advanced Server & Workload Protection (4)"
	sapScannerForVisionOneEndpointSecurityPro creditAllocation = "SAP Scanner for Vision One - Endpoint Security (Pro) (8)"
)

func (c creditAllocation) String() string {
	return string(c)
}
