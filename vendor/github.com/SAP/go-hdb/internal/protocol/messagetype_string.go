// Code generated by "stringer -type=messageType"; DO NOT EDIT.

package protocol

import "strconv"

const (
	_messageType_name_0 = "mtNil"
	_messageType_name_1 = "mtExecuteDirectmtPreparemtAbapStreammtXAStartmtXAJoin"
	_messageType_name_2 = "mtExecute"
	_messageType_name_3 = "mtWriteLobmtReadLobmtFindLob"
	_messageType_name_4 = "mtAuthenticatemtConnectmtCommitmtRollbackmtCloseResultsetmtDropStatementIDmtFetchNextmtFetchAbsolutemtFetchRelativemtFetchFirstmtFetchLast"
	_messageType_name_5 = "mtDisconnectmtExecuteITabmtFetchNextITabmtInsertNextITab"
)

var (
	_messageType_index_1 = [...]uint8{0, 15, 24, 36, 45, 53}
	_messageType_index_3 = [...]uint8{0, 10, 19, 28}
	_messageType_index_4 = [...]uint8{0, 14, 23, 31, 41, 57, 74, 85, 100, 115, 127, 138}
	_messageType_index_5 = [...]uint8{0, 12, 25, 40, 56}
)

func (i messageType) String() string {
	switch {
	case i == 0:
		return _messageType_name_0
	case 2 <= i && i <= 6:
		i -= 2
		return _messageType_name_1[_messageType_index_1[i]:_messageType_index_1[i+1]]
	case i == 13:
		return _messageType_name_2
	case 16 <= i && i <= 18:
		i -= 16
		return _messageType_name_3[_messageType_index_3[i]:_messageType_index_3[i+1]]
	case 65 <= i && i <= 75:
		i -= 65
		return _messageType_name_4[_messageType_index_4[i]:_messageType_index_4[i+1]]
	case 77 <= i && i <= 80:
		i -= 77
		return _messageType_name_5[_messageType_index_5[i]:_messageType_index_5[i+1]]
	default:
		return "messageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
