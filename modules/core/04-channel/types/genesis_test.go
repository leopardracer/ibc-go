package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v10/modules/core/24-host"
	ibcerrors "github.com/cosmos/ibc-go/v10/modules/core/errors"
)

const (
	testPort1         = "firstport"
	testPort2         = "secondport"
	testConnectionIDA = "connectionidatob"

	testChannel1 = "channel-0"
	testChannel2 = "channel-1"

	testChannelOrder   = types.ORDERED
	testChannelVersion = "1.0"
)

func TestValidateGenesis(t *testing.T) {
	counterparty1 := types.NewCounterparty(testPort1, testChannel1)
	counterparty2 := types.NewCounterparty(testPort2, testChannel2)
	testCases := []struct {
		name     string
		genState types.GenesisState
		expErr   error
	}{
		{
			name:     "default",
			genState: types.DefaultGenesisState(),
			expErr:   nil,
		},
		{
			name: "valid genesis",
			genState: types.NewGenesisState(
				[]types.IdentifiedChannel{
					types.NewIdentifiedChannel(
						testPort1, testChannel1, types.NewChannel(
							types.INIT, testChannelOrder, counterparty2, []string{testConnectionIDA}, testChannelVersion,
						),
					),
					types.NewIdentifiedChannel(
						testPort2, testChannel2, types.NewChannel(
							types.INIT, testChannelOrder, counterparty1, []string{testConnectionIDA}, testChannelVersion,
						),
					),
				},
				[]types.PacketState{
					types.NewPacketState(testPort2, testChannel2, 1, []byte("ack")),
				},
				[]types.PacketState{
					types.NewPacketState(testPort2, testChannel2, 1, []byte("")),
				},
				[]types.PacketState{
					types.NewPacketState(testPort1, testChannel1, 1, []byte("commit_hash")),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort1, testChannel1, 1),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort2, testChannel2, 1),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort2, testChannel2, 1),
				},
				2,
			),
			expErr: nil,
		},
		{
			name: "invalid channel",
			genState: types.GenesisState{
				Channels: []types.IdentifiedChannel{
					types.NewIdentifiedChannel(
						testPort1, "(testChannel1)", types.NewChannel(
							types.INIT, testChannelOrder, counterparty2, []string{testConnectionIDA}, testChannelVersion,
						),
					),
				},
			},
			expErr: host.ErrInvalidID,
		},
		{
			name: "invalid ack",
			genState: types.GenesisState{
				Acknowledgements: []types.PacketState{
					types.NewPacketState(testPort2, testChannel2, 1, nil),
				},
			},
			expErr: types.ErrInvalidAcknowledgement,
		},
		{
			name: "invalid commitment",
			genState: types.GenesisState{
				Commitments: []types.PacketState{
					types.NewPacketState(testPort1, testChannel1, 1, nil),
				},
			},
			expErr: types.ErrInvalidCommitment,
		},
		{
			name: "invalid send seq",
			genState: types.GenesisState{
				SendSequences: []types.PacketSequence{
					types.NewPacketSequence(testPort1, testChannel1, 0),
				},
			},
			expErr: ibcerrors.ErrInvalidSequence,
		},
		{
			name: "invalid recv seq",
			genState: types.GenesisState{
				RecvSequences: []types.PacketSequence{
					types.NewPacketSequence(testPort1, "(testChannel1)", 1),
				},
			},
			expErr: host.ErrInvalidID,
		},
		{
			name: "invalid recv seq 2",
			genState: types.GenesisState{
				RecvSequences: []types.PacketSequence{
					types.NewPacketSequence("(testPort1)", testChannel1, 1),
				},
			},
			expErr: host.ErrInvalidID,
		},
		{
			name: "invalid ack seq",
			genState: types.GenesisState{
				AckSequences: []types.PacketSequence{
					types.NewPacketSequence(testPort1, "(testChannel1)", 1),
				},
			},
			expErr: types.ErrInvalidAcknowledgement,
		},
		{
			name: "invalid channel identifier",
			genState: types.NewGenesisState(
				[]types.IdentifiedChannel{
					types.NewIdentifiedChannel(
						testPort1, "chan-0", types.NewChannel(
							types.INIT, testChannelOrder, counterparty2, []string{testConnectionIDA}, testChannelVersion,
						),
					),
					types.NewIdentifiedChannel(
						testPort2, testChannel2, types.NewChannel(
							types.INIT, testChannelOrder, counterparty1, []string{testConnectionIDA}, testChannelVersion,
						),
					),
				},
				[]types.PacketState{
					types.NewPacketState(testPort2, testChannel2, 1, []byte("ack")),
				},
				[]types.PacketState{
					types.NewPacketState(testPort2, testChannel2, 1, []byte("")),
				},
				[]types.PacketState{
					types.NewPacketState(testPort1, testChannel1, 1, []byte("commit_hash")),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort1, testChannel1, 1),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort2, testChannel2, 1),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort2, testChannel2, 1),
				},
				0,
			),
			expErr: host.ErrInvalidID,
		},
		{
			name: "next channel sequence is less than maximum channel identifier sequence used",
			genState: types.NewGenesisState(
				[]types.IdentifiedChannel{
					types.NewIdentifiedChannel(
						testPort1, "channel-10", types.NewChannel(
							types.INIT, testChannelOrder, counterparty2, []string{testConnectionIDA}, testChannelVersion,
						),
					),
					types.NewIdentifiedChannel(
						testPort2, testChannel2, types.NewChannel(
							types.INIT, testChannelOrder, counterparty1, []string{testConnectionIDA}, testChannelVersion,
						),
					),
				},
				[]types.PacketState{
					types.NewPacketState(testPort2, testChannel2, 1, []byte("ack")),
				},
				[]types.PacketState{
					types.NewPacketState(testPort2, testChannel2, 1, []byte("")),
				},
				[]types.PacketState{
					types.NewPacketState(testPort1, testChannel1, 1, []byte("commit_hash")),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort1, testChannel1, 1),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort2, testChannel2, 1),
				},
				[]types.PacketSequence{
					types.NewPacketSequence(testPort2, testChannel2, 1),
				},
				0,
			),
			expErr: ibcerrors.ErrInvalidSequence,
		},
	}

	for _, tc := range testCases {
		err := tc.genState.Validate()
		if tc.expErr == nil {
			require.NoError(t, err, tc.name)
		} else {
			require.Error(t, err, tc.name)
			require.ErrorIs(t, err, tc.expErr)
		}
	}
}
