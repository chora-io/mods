package genesis

import (
	"encoding/json"

	"cosmossdk.io/orm/model/ormdb"
	"google.golang.org/protobuf/proto"

	"github.com/chora-io/mods/voucher"
	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
	v1 "github.com/chora-io/mods/voucher/types/v1"
	"github.com/chora-io/mods/voucher/utils"
)

// ValidateGenesis validates genesis state.
func ValidateGenesis(bz json.RawMessage) error {
	err := validateJSON(bz)
	if err != nil {
		return err
	}

	return nil
}

func validateJSON(bz json.RawMessage) error {
	_, err := ormdb.NewModuleDB(&voucher.ModuleSchema, ormdb.ModuleDBOptions{
		JSONValidator: validateMsg,
	})
	if err != nil {
		return err
	}

	//src, err := ormjson.NewRawMessageSource(bz)
	//if err != nil {
	//	return err
	//}
	//
	//err = db.ValidateJSON(src)
	//if err != nil {
	//	return err
	//}

	return nil
}

func validateMsg(msg proto.Message) error {
	switch msg.(type) {

	case *voucherv1.Voucher:
		m := &v1.Voucher{}
		if err := utils.PulsarToGogoSlow(msg, m); err != nil {
			return err
		}
		return m.Validate()

	case *voucherv1.Balance:
		m := &v1.Balance{}
		if err := utils.PulsarToGogoSlow(msg, m); err != nil {
			return err
		}
		return m.Validate()
	}

	return nil
}
