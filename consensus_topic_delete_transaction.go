package hedera

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

type ConsensusTopicDeleteTransaction struct {
	TransactionBuilder
	pb *proto.ConsensusDeleteTopicTransactionBody
}

func NewConsensusTopicDeleteTransaction() ConsensusTopicDeleteTransaction {
	pb := &proto.ConsensusDeleteTopicTransactionBody{}

	inner := newTransactionBuilder()
	inner.pb.Data = &proto.TransactionBody_ConsensusDeleteTopic{pb}

	builder := ConsensusTopicDeleteTransaction{inner, pb}

	return builder
}

func (builder ConsensusTopicDeleteTransaction) SetTopicID(id ConsensusTopicID) ConsensusTopicDeleteTransaction {
	builder.pb.TopicID = id.toProto()
	return builder
}

func (builder ConsensusTopicDeleteTransaction) Build(client *Client) (Transaction, error) {
	return builder.TransactionBuilder.Build(client)
}

//
// The following _5_ must be copy-pasted at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func (builder ConsensusTopicDeleteTransaction) SetMaxTransactionFee(maxTransactionFee Hbar) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetMaxTransactionFee(maxTransactionFee), builder.pb}
}

func (builder ConsensusTopicDeleteTransaction) SetMemo(memo string) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetTransactionMemo(memo), builder.pb}
}

func (builder ConsensusTopicDeleteTransaction) SetTransactionValidDuration(validDuration time.Duration) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetTransactionValidDuration(validDuration), builder.pb}
}

func (builder ConsensusTopicDeleteTransaction) SetTransactionID(transactionID TransactionID) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetTransactionID(transactionID), builder.pb}
}

func (builder ConsensusTopicDeleteTransaction) SetNodeAccountID(nodeAccountID AccountID) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetNodeAccountID(nodeAccountID), builder.pb}
}
