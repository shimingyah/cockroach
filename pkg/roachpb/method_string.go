// Code generated by "stringer -type=Method"; DO NOT EDIT.

package roachpb

import "fmt"

const _Method_name = "GetPutConditionalPutIncrementDeleteDeleteRangeClearRangeScanReverseScanBeginTransactionEndTransactionAdminSplitAdminMergeAdminTransferLeaseAdminChangeReplicasHeartbeatTxnGCPushTxnQueryTxnRangeLookupResolveIntentResolveIntentRangeNoopMergeTruncateLogRequestLeaseTransferLeaseLeaseInfoComputeChecksumDeprecatedVerifyChecksumCheckConsistencyInitPutWriteBatchExportImportAdminScatterAddSSTable"

var _Method_index = [...]uint16{0, 3, 6, 20, 29, 35, 46, 56, 60, 71, 87, 101, 111, 121, 139, 158, 170, 172, 179, 187, 198, 211, 229, 233, 238, 249, 261, 274, 283, 298, 322, 338, 345, 355, 361, 367, 379, 389}

func (i Method) String() string {
	if i < 0 || i >= Method(len(_Method_index)-1) {
		return fmt.Sprintf("Method(%d)", i)
	}
	return _Method_name[_Method_index[i]:_Method_index[i+1]]
}
