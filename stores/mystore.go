// Copyright 2016-2018 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stores

import (
	"os"
	"path/filepath"

	"github.com/nats-io/go-nats-streaming/pb"
	"github.com/nats-io/nats-streaming-server/spb"
)

//MyStore a fast message test
type MyStore struct {
}

//GetExclusiveLock implements Store Interface
func (ms *MyStore) GetExclusiveLock() (bool, error) {
	return true, nil
}

// Init can be used to initialize the store with server's information.
func (ms *MyStore) Init(info *spb.ServerInfo) error {
	return nil
}

// // Name returns the name type of this store (e.g: MEMORY, FILESTORE, etc...).
// func (ms *MyStore) Name() string {

// }

// Recover returns the recovered state.
// Implementations that do not persist state and therefore cannot
// recover from a previous run MUST return nil, not an error.
// However, an error must be returned for implementations that are
// attempting to recover the state but fail to do so.
func (ms *MyStore) Recover() (*RecoveredState, error) {
	return nil, nil
}

// SetLimits sets limits for this store. The action is not expected
// to be retroactive.
// The store implementation should make a deep copy as to not change
// the content of the structure passed by the caller.
func (ms *MyStore) SetLimits(limits *StoreLimits) error {
	return nil
}

// GetChannelLimits returns the limit for this channel. If the channel
// does not exist, returns nil.
func (ms *MyStore) GetChannelLimits(name string) *ChannelLimits {
	return nil
}

// CreateChannel creates a Channel.
// Implementations should return ErrAlreadyExists if the channel was
// already created.
// Limits defined for this channel in StoreLimits.PeChannel map, if present,
// will apply. Otherwise, the global limits in StoreLimits will apply.
func (ms *MyStore) CreateChannel(channel string) (*Channel, error) {

	path := filepath.Join("./", channel)
	if _, err := os.Stat(path); os.IsExist(err) {
		return nil, ErrAlreadyExists
	}

	os.Mkdir(path, os.ModeDir+os.ModePerm)
	subs := &MySubStore{}

	msgs := &MyMsgStore{}

	c := &Channel{
		Subs: subs,
		Msgs: msgs,
	}

	return c, nil
}

// DeleteChannel deletes a Channel.
// Implementations should make sure that if no error is returned, the
// channel would not be recovered after a restart, unless CreateChannel()
// with the same channel is invoked.
// If processing is expecting to be time consuming, work should be done
// in the background as long as the above condition is guaranteed.
// It is also acceptable for an implementation to have CreateChannel()
// return an error if background deletion is still happening for a
// channel of the same name.
func (ms *MyStore) DeleteChannel(channel string) error {
	path := filepath.Join("./", channel)
	if _, err := os.Stat(path); os.IsExist(err) {

	}

	return nil
}

// AddClient stores information about the client identified by `clientID`.
func (ms *MyStore) AddClient(info *spb.ClientInfo) (*Client, error) {

}

// DeleteClient removes the client identified by `clientID` from the store.
func (ms *MyStore) DeleteClient(clientID string) error {
	return nil
}

// Close closes this store (including all MsgStore and SubStore).
// If an exclusive lock was acquired, the lock shall be released.
func (ms *MyStore) Close() error {
	return nil
}

// SubStore is the interface for storage of Subscriptions on a given channel.
//
// Implementations of this interface should not attempt to validate that
// a subscription is valid (that is, has not been deleted) when processing
// updates.
type MySubStore struct{}

// CreateSub records a new subscription represented by SubState. On success,
// it records the subscription's ID in SubState.ID. This ID is to be used
// by the other SubStore methods.
func (mss *MySubStore) CreateSub(*spb.SubState) error {

}

// UpdateSub updates a given subscription represented by SubState.
func (mss *MySubStore) UpdateSub(*spb.SubState) error {

}

// DeleteSub invalidates the subscription 'subid'.
func (mss *MySubStore) DeleteSub(subid uint64) error {

}

// AddSeqPending adds the given message 'seqno' to the subscription 'subid'.
func (mss *MySubStore) AddSeqPending(subid, seqno uint64) error {

}

// AckSeqPending records that the given message 'seqno' has been acknowledged
// by the subscription 'subid'.
func (mss *MySubStore) AckSeqPending(subid, seqno uint64) error {

}

// Flush is for stores that may buffer operations and need them to be persisted.
func (mss *MySubStore) Flush() error {

}

// Close closes the subscriptions store.
func (mss *MySubStore) Close() error {

}

// MsgStore is the interface for storage of Messages on a given channel.
type MyMsgStore struct{}

// State returns some statistics related to this store.
func (mms *MyMsgStore) State() (numMessages int, byteSize uint64, err error) {

}

// Store stores a message and returns the message sequence.
func (mms *MyMsgStore) Store(msg *pb.MsgProto) (uint64, error) {

}

// Lookup returns the stored message with given sequence number.
func (mms *MyMsgStore) Lookup(seq uint64) (*pb.MsgProto, error) {

}

// FirstSequence returns sequence for first message stored, 0 if no
// message is stored.
func (mms *MyMsgStore) FirstSequence() (uint64, error) {

}

// LastSequence returns sequence for last message stored, 0 if no
// message is stored.
func (mms *MyMsgStore) LastSequence() (uint64, error) {

}

// FirstAndLastSequence returns sequences for the first and last messages stored,
// 0 if no message is stored.
func (mms *MyMsgStore) FirstAndLastSequence() (uint64, uint64, error) {

}

// GetSequenceFromTimestamp returns the sequence of the first message whose
// timestamp is greater or equal to given timestamp.
func (mms *MyMsgStore) GetSequenceFromTimestamp(timestamp int64) (uint64, error) {

}

// FirstMsg returns the first message stored.
func (mms *MyMsgStore) FirstMsg() (*pb.MsgProto, error) {

}

// LastMsg returns the last message stored.
func (mms *MyMsgStore) LastMsg() (*pb.MsgProto, error) {

}

// Flush is for stores that may buffer operations and need them to be persisted.
func (mms *MyMsgStore) Flush() error {

}

// Empty removes all messages from the store
func (mms *MyMsgStore) Empty() error {

}

// Close closes the store.
func (mms *MyMsgStore) Close() error {

}
