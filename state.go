package substrate

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/centrifuge/go-substrate-rpc-client/scale"
)

// MethodIDX [sectionIndex, methodIndex] 16bits
type MethodIDX struct {
	SectionIndex uint8
	MethodIndex uint8
}

func (e *MethodIDX) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&e.SectionIndex)
	if err != nil {
		return err
	}

	err = decoder.Decode(&e.MethodIndex)
	if err != nil {
		return err
	}

	return nil
}

func (m MethodIDX) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.SectionIndex)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.MethodIndex)
	if err != nil {
		return err
	}

	return nil
}

type MetadataV4 struct {
	Modules []ModuleMetaData
}

func (m *MetadataV4) MethodIndex(method string) MethodIDX {
	s := strings.Split(method, ".")
	var sIDX, mIDX uint8 = 0, 0
	// section index
	var sCounter = 0

	for _, n := range m.Modules {
		if n.CallsOptional == 1 {
			if n.Name == s[0] {
				sIDX = uint8(sCounter)
				for j, f := range n.Calls {
					if f.Name == s[1] {
						mIDX = uint8(j)
					}
				}
			}
			sCounter++
		}
	}

	return MethodIDX{sIDX, mIDX}
}

func (m *MetadataV4) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Modules)
	if err != nil {
		return err
	}
	return nil
}

type FunctionArgumentMetadata struct {
	Name string
	Type string
}

func (m *FunctionArgumentMetadata) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Type)
	if err != nil {
		return err
	}

	return nil
}

type FunctionMetaData struct {
	Name string
	Args []FunctionArgumentMetadata
	Documentation []string
}

func (m *FunctionMetaData) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Args)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Documentation)
	if err != nil {
		return err
	}

	return nil
}

type EventMetadata struct {
	Name string
	Args []string
	Documentation []string
}

func (m *EventMetadata) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Args)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Documentation)
	if err != nil {
		return err
	}

	return nil
}

/**
[{"name":"AccountNonce","modifier":"Default","type":{"MapType":{"hasher":"Blake2_256","key":"AccountId","value":"Index","isLinked":false}},"fallback":"0x0000000000000000","documentation":[" Extrinsics nonce for accounts."]},{"name":"ExtrinsicCount","modifier":"Optional","type":{"PlainType":"u32"},"fallback":"0x00","documentation":[" Total extrinsics count for the current block."]},{"name":"AllExtrinsicsLen","modifier":"Optional","type":{"PlainType":"u32"},"fallback":"0x00","documentation":[" Total length in bytes for all extrinsics put together, for the current block."]},{"name":"BlockHash","modifier":"Default","type":{"MapType":{"hasher":"Blake2_256","key":"BlockNumber","value":"Hash","isLinked":false}},"fallback":"0x0000000000000000000000000000000000000000000000000000000000000000","documentation":[" Map of block numbers to block hashes."]},{"name":"ExtrinsicData","modifier":"Default","type":{"MapType":{"hasher":"Blake2_256","key":"u32","value":"Bytes","isLinked":false}},"fallback":"0x00","documentation":[" Extrinsics data for the current block (maps extrinsic's index to its data)."]},{"name":"RandomSeed","modifier":"Default","type":{"PlainType":"Hash"},"fallback":"0x0000000000000000000000000000000000000000000000000000000000000000","documentation":[" Random seed of the current block."]},{"name":"Number","modifier":"Default","type":{"PlainType":"BlockNumber"},"fallback":"0x0000000000000000","documentation":[" The current block number being processed. Set by `execute_block`."]},{"name":"ParentHash","modifier":"Default","type":{"PlainType":"Hash"},"fallback":"0x0000000000000000000000000000000000000000000000000000000000000000","documentation":[" Hash of the previous block."]},{"name":"ExtrinsicsRoot","modifier":"Default","type":{"PlainType":"Hash"},"fallback":"0x0000000000000000000000000000000000000000000000000000000000000000","documentation":[" Extrinsics root of the current block, also part of the block header."]},{"name":"Digest","modifier":"Default","type":{"PlainType":"Digest"},"fallback":"0x00","documentation":[" Digest of the current block, also part of the block header."]},{"name":"Events","modifier":"Default","type":{"PlainType":"Vec<EventRecord>"},"fallback":"0x00","documentation":[" Events deposited for the current block."]}]
 */

type TypMap struct {
	Hasher uint8
	Key string
	Value string
	IsLinked bool
}

func (m *TypMap) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Hasher)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Key)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Value)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.IsLinked)
	if err != nil {
		return err
	}

	return nil
}

type TypDoubleMap struct {
	Hasher uint8
	Key string
	Key2 string
	Value string
	Key2Hasher string
}

func (m *TypDoubleMap) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Hasher)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Key)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Key2)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Value)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Key2Hasher)
	if err != nil {
		return err
	}

	return nil
}

type StorageFunctionMetadata struct {
	Name string
	Modifier uint8
	Type uint8
	Plane string
	Map TypMap
	DMap TypDoubleMap
	Fallback []byte
	Documentation []string
}

func (m *StorageFunctionMetadata) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Modifier)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Type)
	if err != nil {
		return err
	}

	switch m.Type {
	case 0:
		err = decoder.Decode(&m.Plane)
		if err != nil {
			return err
		}
	case 1:
		err = decoder.Decode(&m.Map)
		if err != nil {
			return err
		}
	default:
		err = decoder.Decode(&m.DMap)
		if err != nil {
			return err
		}
	}
	err = decoder.Decode(&m.Fallback)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Documentation)
	if err != nil {
		return err
	}
	// fmt.Println(m.Documentation)
	return nil
}

type ModuleMetaData struct {
	Name string
	Prefix string
	StorageOptional uint8
	Storage []StorageFunctionMetadata
	CallsOptional uint8
	Calls []FunctionMetaData
	EventsOptional uint8
	Events []EventMetadata
}

func (m *ModuleMetaData) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Prefix)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.StorageOptional)
	if err != nil {
		return err
	}

	if m.StorageOptional == 1 {
		err = decoder.Decode(&m.Storage)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.CallsOptional)
	if err != nil {
		return err
	}

	if m.CallsOptional == 1 {
		err = decoder.Decode(&m.Calls)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.EventsOptional)
	if err != nil {
		return err
	}

	if m.EventsOptional == 1 {
		err = decoder.Decode(&m.Events)
		if err != nil {
			return err
		}
		// fmt.Println(m.Events)
	}
	return nil
}

// MetadataVersioned only supports v4
type MetadataVersioned struct {
	// 1635018093
	MagicNumber   uint32
	Version       uint8
	Metadata      MetadataV4
}

func NewMetadataVersioned() *MetadataVersioned {
	return &MetadataVersioned{Metadata:MetadataV4{make([]ModuleMetaData, 0)}}
}

func (m *MetadataVersioned) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.MagicNumber)
	if err != nil {
		return err
	}
	// we need to decide which struct to use based on the following number(enum), for now its hardcoded
	err = decoder.Decode(&m.Version)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Metadata)
	if err != nil {
		return err
	}

	return nil
}

type State struct {
	nonetwork bool
	client Client
} 

func NewStateRPC(client Client) *State {
	return &State{client:client}
}

func (s *State) MetaData(blockHash Hash) (*MetadataVersioned, error) {
	// "0xd133045f0efad58582772cbdb6f5f0cd6af7bb4bf1f30d039a4b18b4bdaf4901"
	var res string
	if !s.nonetwork {
		// block hash can give error - Error(Client(UnknownBlock("State already discarded for Hash(0xxxx)")), State { next_error: None, backtrace: InternalBacktrace { backtrace: None } })
		err := s.client.Call(&res, "state_getMetadata", blockHash.String())
		if err != nil {
			return nil, err
		}
	} else {
		// just for testing
		res = "0x6d65746104201873797374656d1853797374656d012c304163636f756e744e6f6e636501010130543a3a4163636f756e74496420543a3a496e64657800200000000000000000047c2045787472696e73696373206e6f6e636520666f72206163636f756e74732e3845787472696e736963436f756e7400000c753332040004b820546f74616c2065787472696e7369637320636f756e7420666f72207468652063757272656e7420626c6f636b2e40416c6c45787472696e736963734c656e00000c753332040004390120546f74616c206c656e67746820696e20627974657320666f7220616c6c2065787472696e736963732070757420746f6765746865722c20666f72207468652063757272656e7420626c6f636b2e24426c6f636b4861736801010138543a3a426c6f636b4e756d6265721c543a3a48617368008000000000000000000000000000000000000000000000000000000000000000000498204d6170206f6620626c6f636b206e756d6265727320746f20626c6f636b206861736865732e3445787472696e736963446174610101010c7533321c5665633c75383e0004000431012045787472696e73696373206461746120666f72207468652063757272656e7420626c6f636b20286d6170732065787472696e736963277320696e64657820746f206974732064617461292e2852616e646f6d5365656401001c543a3a4861736880000000000000000000000000000000000000000000000000000000000000000004882052616e646f6d2073656564206f66207468652063757272656e7420626c6f636b2e184e756d626572010038543a3a426c6f636b4e756d626572200000000000000000040901205468652063757272656e7420626c6f636b206e756d626572206265696e672070726f6365737365642e205365742062792060657865637574655f626c6f636b602e28506172656e744861736801001c543a3a4861736880000000000000000000000000000000000000000000000000000000000000000004702048617368206f66207468652070726576696f757320626c6f636b2e3845787472696e73696373526f6f7401001c543a3a486173688000000000000000000000000000000000000000000000000000000000000000000415012045787472696e7369637320726f6f74206f66207468652063757272656e7420626c6f636b2c20616c736f2070617274206f662074686520626c6f636b206865616465722e18446967657374010024543a3a446967657374040004f020446967657374206f66207468652063757272656e7420626c6f636b2c20616c736f2070617274206f662074686520626c6f636b206865616465722e184576656e74730100685665633c4576656e745265636f72643c543a3a4576656e743e3e040004a0204576656e7473206465706f736974656420666f72207468652063757272656e7420626c6f636b2e0001084045787472696e7369635375636365737300049420416e2065787472696e73696320636f6d706c65746564207375636365737366756c6c792e3c45787472696e7369634661696c656400045420416e2065787472696e736963206661696c65642e2474696d657374616d702454696d657374616d7001100c4e6f77010024543a3a4d6f6d656e7420000000000000000004902043757272656e742074696d6520666f72207468652063757272656e7420626c6f636b2e2c426c6f636b506572696f64000024543a3a4d6f6d656e740400044501204f6c642073746f72616765206974656d2070726f766964656420666f7220636f6d7061746962696c6974792e2052656d6f766520616674657220616c6c206e6574776f726b732075706772616465642e344d696e696d756d506572696f64010024543a3a4d6f6d656e7420030000000000000010690120546865206d696e696d756d20706572696f64206265747765656e20626c6f636b732e204265776172652074686174207468697320697320646966666572656e7420746f20746865202a65787065637465642a20706572696f64690120746861742074686520626c6f636b2070726f64756374696f6e206170706172617475732070726f76696465732e20596f75722063686f73656e20636f6e73656e7375732073797374656d2077696c6c2067656e6572616c6c79650120776f726b2077697468207468697320746f2064657465726d696e6520612073656e7369626c6520626c6f636b2074696d652e20652e672e20466f7220417572612c2069742077696c6c20626520646f75626c6520746869737020706572696f64206f6e2064656661756c742073657474696e67732e24446964557064617465010010626f6f6c040004b420446964207468652074696d657374616d7020676574207570646174656420696e207468697320626c6f636b3f01040c736574040c6e6f7748436f6d706163743c543a3a4d6f6d656e743e205820536574207468652063757272656e742074696d652e00750120546869732063616c6c2073686f756c6420626520696e766f6b65642065786163746c79206f6e63652070657220626c6f636b2e2049742077696c6c2070616e6963206174207468652066696e616c697a6174696f6e2070686173652cbc20696620746869732063616c6c206861736e2774206265656e20696e766f6b656420627920746861742074696d652e008d01205468652074696d657374616d702073686f756c642062652067726561746572207468616e207468652070726576696f7573206f6e652062792074686520616d6f756e742073706563696669656420627920606d696e696d756d5f706572696f64602e00d820546865206469737061746368206f726967696e20666f7220746869732063616c6c206d7573742062652060496e686572656e74602e0024636f6e73656e73757324436f6e73656e73757301044c4f726967696e616c417574686f7269746965730000485665633c543a3a53657373696f6e4b65793e040000011c487265706f72745f6d69736265686176696f72041c5f7265706f72741c5665633c75383e0464205265706f727420736f6d65206d69736265686176696f722e306e6f74655f6f66666c696e65041c6f66666c696e65f43c543a3a496e686572656e744f66666c696e655265706f727420617320496e686572656e744f66666c696e655265706f72743e3a3a496e686572656e74045101204e6f74652074686174207468652070726576696f757320626c6f636b27732076616c696461746f72206d697373656420697473206f70706f7274756e69747920746f2070726f706f7365206120626c6f636b2e1872656d61726b041c5f72656d61726b1c5665633c75383e046c204d616b6520736f6d65206f6e2d636861696e2072656d61726b2e387365745f686561705f7061676573041470616765730c75363404fc2053657420746865206e756d626572206f6620706167657320696e2074686520576562417373656d626c7920656e7669726f6e6d656e74277320686561702e207365745f636f6465040c6e65771c5665633c75383e04482053657420746865206e657720636f64652e2c7365745f73746f7261676504146974656d73345665633c4b657956616c75653e046c2053657420736f6d65206974656d73206f662073746f726167652e306b696c6c5f73746f7261676504106b657973205665633c4b65793e0478204b696c6c20736f6d65206974656d732066726f6d2073746f726167652e001061757261000000001c696e64696365731c496e646963657301082c4e657874456e756d53657401003c543a3a4163636f756e74496e6465781000000000047c20546865206e657874206672656520656e756d65726174696f6e207365742e1c456e756d5365740101013c543a3a4163636f756e74496e646578445665633c543a3a4163636f756e7449643e00040004582054686520656e756d65726174696f6e20736574732e010001043c4e65774163636f756e74496e64657808244163636f756e744964304163636f756e74496e64657810882041206e6577206163636f756e7420696e646578207761732061737369676e65642e0005012054686973206576656e74206973206e6f7420747269676765726564207768656e20616e206578697374696e6720696e64657820697320726561737369676e65646020746f20616e6f7468657220604163636f756e744964602e2062616c616e6365732042616c616e636573012834546f74616c49737375616e6365010028543a3a42616c616e6365400000000000000000000000000000000004982054686520746f74616c20756e6974732069737375656420696e207468652073797374656d2e484578697374656e7469616c4465706f736974010028543a3a42616c616e6365400000000000000000000000000000000004d420546865206d696e696d756d20616d6f756e7420726571756972656420746f206b65657020616e206163636f756e74206f70656e2e2c5472616e73666572466565010028543a3a42616c616e636540000000000000000000000000000000000494205468652066656520726571756972656420746f206d616b652061207472616e736665722e2c4372656174696f6e466565010028543a3a42616c616e63654000000000000000000000000000000000049c205468652066656520726571756972656420746f2063726561746520616e206163636f756e742e485472616e73616374696f6e42617365466565010028543a3a42616c616e6365400000000000000000000000000000000004dc205468652066656520746f206265207061696420666f72206d616b696e672061207472616e73616374696f6e3b2074686520626173652e485472616e73616374696f6e42797465466565010028543a3a42616c616e63654000000000000000000000000000000000040d01205468652066656520746f206265207061696420666f72206d616b696e672061207472616e73616374696f6e3b20746865207065722d6279746520706f7274696f6e2e1c56657374696e6700010130543a3a4163636f756e7449646c56657374696e675363686564756c653c543a3a42616c616e63653e00040004d820496e666f726d6174696f6e20726567617264696e67207468652076657374696e67206f66206120676976656e206163636f756e742e2c4672656542616c616e636501010130543a3a4163636f756e74496428543a3a42616c616e63650040000000000000000000000000000000002c9c20546865202766726565272062616c616e6365206f66206120676976656e206163636f756e742e004101205468697320697320746865206f6e6c792062616c616e63652074686174206d61747465727320696e207465726d73206f66206d6f7374206f7065726174696f6e73206f6e20746f6b656e732e204974750120616c6f6e65206973207573656420746f2064657465726d696e65207468652062616c616e6365207768656e20696e2074686520636f6e747261637420657865637574696f6e20656e7669726f6e6d656e742e205768656e207468697355012062616c616e63652066616c6c732062656c6f77207468652076616c7565206f6620604578697374656e7469616c4465706f736974602c207468656e20746865202763757272656e74206163636f756e74272069733d012064656c657465643a207370656369666963616c6c7920604672656542616c616e6365602e20467572746865722c2074686520604f6e4672656542616c616e63655a65726f602063616c6c6261636b450120697320696e766f6b65642c20676976696e672061206368616e636520746f2065787465726e616c206d6f64756c657320746f20636c65616e2075702064617461206173736f636961746564207769746854207468652064656c65746564206163636f756e742e005d01206073797374656d3a3a4163636f756e744e6f6e63656020697320616c736f2064656c657465642069662060526573657276656442616c616e63656020697320616c736f207a65726f2028697420616c736f2067657473150120636f6c6c617073656420746f207a65726f2069662069742065766572206265636f6d6573206c657373207468616e20604578697374656e7469616c4465706f736974602e3c526573657276656442616c616e636501010130543a3a4163636f756e74496428543a3a42616c616e63650040000000000000000000000000000000002c75012054686520616d6f756e74206f66207468652062616c616e6365206f66206120676976656e206163636f756e7420746861742069732065787465726e616c6c792072657365727665643b20746869732063616e207374696c6c206765749c20736c61736865642c20627574206765747320736c6173686564206c617374206f6620616c6c2e006d0120546869732062616c616e63652069732061202772657365727665272062616c616e63652074686174206f746865722073756273797374656d732075736520696e206f7264657220746f2073657420617369646520746f6b656e732501207468617420617265207374696c6c20276f776e65642720627920746865206163636f756e7420686f6c6465722c20627574207768696368206172652073757370656e6461626c652e007501205768656e20746869732062616c616e63652066616c6c732062656c6f77207468652076616c7565206f6620604578697374656e7469616c4465706f736974602c207468656e2074686973202772657365727665206163636f756e7427b42069732064656c657465643a207370656369666963616c6c792c2060526573657276656442616c616e6365602e004d01206073797374656d3a3a4163636f756e744e6f6e63656020697320616c736f2064656c6574656420696620604672656542616c616e63656020697320616c736f207a65726f2028697420616c736f2067657473190120636f6c6c617073656420746f207a65726f2069662069742065766572206265636f6d6573206c657373207468616e20604578697374656e7469616c4465706f736974602e29144c6f636b7301010130543a3a4163636f756e744964b05665633c42616c616e63654c6f636b3c543a3a42616c616e63652c20543a3a426c6f636b4e756d6265723e3e00040004b820416e79206c6971756964697479206c6f636b73206f6e20736f6d65206163636f756e742062616c616e6365732e0108207472616e736665720810646573748c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f757263651476616c75654c436f6d706163743c543a3a42616c616e63653e20d8205472616e7366657220736f6d65206c697175696420667265652062616c616e636520746f20616e6f74686572206163636f756e742e00090120607472616e73666572602077696c6c207365742074686520604672656542616c616e636560206f66207468652073656e64657220616e642072656365697665722e21012049742077696c6c2064656372656173652074686520746f74616c2069737375616e6365206f66207468652073797374656d2062792074686520605472616e73666572466565602e1501204966207468652073656e6465722773206163636f756e742069732062656c6f7720746865206578697374656e7469616c206465706f736974206173206120726573756c74b4206f6620746865207472616e736665722c20746865206163636f756e742077696c6c206265207265617065642e00190120546865206469737061746368206f726967696e20666f7220746869732063616c6c206d75737420626520605369676e65646020627920746865207472616e736163746f722e2c7365745f62616c616e63650c0c77686f8c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f7572636510667265654c436f6d706163743c543a3a42616c616e63653e2072657365727665644c436f6d706163743c543a3a42616c616e63653e209420536574207468652062616c616e636573206f66206120676976656e206163636f756e742e00010120546869732077696c6c20616c74657220604672656542616c616e63656020616e642060526573657276656442616c616e63656020696e2073746f726167652e190120496620746865206e65772066726565206f722072657365727665642062616c616e63652069732062656c6f7720746865206578697374656e7469616c206465706f7369742c25012069742077696c6c20616c736f2064656372656173652074686520746f74616c2069737375616e6365206f66207468652073797374656d202860546f74616c49737375616e63656029d820616e6420726573657420746865206163636f756e74206e6f6e636520286073797374656d3a3a4163636f756e744e6f6e636560292e00b420546865206469737061746368206f726967696e20666f7220746869732063616c6c2069732060726f6f74602e010c284e65774163636f756e7408244163636f756e7449641c42616c616e6365046c2041206e6577206163636f756e742077617320637265617465642e345265617065644163636f756e7404244163636f756e744964045c20416e206163636f756e7420776173207265617065642e205472616e7366657210244163636f756e744964244163636f756e7449641c42616c616e63651c42616c616e636504b0205472616e7366657220737563636565646564202866726f6d2c20746f2c2076616c75652c2066656573292e107375646f105375646f01040c4b6579010030543a3a4163636f756e74496480000000000000000000000000000000000000000000000000000000000000000004842054686520604163636f756e74496460206f6620746865207375646f206b65792e0108107375646f042070726f706f73616c40426f783c543a3a50726f706f73616c3e0c39012041757468656e7469636174657320746865207375646f206b657920616e64206469737061746368657320612066756e6374696f6e2063616c6c20776974682060526f6f7460206f726967696e2e00d020546865206469737061746368206f726967696e20666f7220746869732063616c6c206d757374206265205f5369676e65645f2e1c7365745f6b6579040c6e65778c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f757263650c75012041757468656e74696361746573207468652063757272656e74207375646f206b657920616e6420736574732074686520676976656e204163636f756e7449642028606e6577602920617320746865206e6577207375646f206b65792e00d020546865206469737061746368206f726967696e20666f7220746869732063616c6c206d757374206265205f5369676e65645f2e01081453756469640410626f6f6c04602041207375646f206a75737420746f6f6b20706c6163652e284b65794368616e67656404244163636f756e74496404f020546865207375646f6572206a757374207377697463686564206964656e746974793b20746865206f6c64206b657920697320737570706c6965642e206b6572706c756e6b204b6572706c756e6b01041c416e63686f72730101011c543a3a486173687c416e63686f723c543a3a486173682c20543a3a426c6f636b4e756d6265723e00210100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010418636f6d6d69740c48616e63686f725f69645f707265696d6167651c543a3a4861736820646f635f726f6f741c543a3a486173681470726f6f661c543a3a486173680001043c416e63686f72436f6d6d697474656410244163636f756e744964104861736810486173682c426c6f636b4e756d62657200"
	}

	b, err := hexutil.Decode(res)
	if err != nil {
		return nil, err
	}

	dec := scale.NewDecoder(bytes.NewReader(b))
	n := NewMetadataVersioned()
	err = dec.Decode(n)
	if err != nil {
		return nil, err
	}

	return n, nil
}

// Keys state_getKeys
func (s *State) Keys(blockHash Hash) (*MetadataVersioned, error) {
	var res string
	if !s.nonetwork {
		err := s.client.Call(&res, "state_getKeys", blockHash.String())
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}