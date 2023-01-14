package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrDBCreate = func(err error) error {
		return errors.Wrap(err, "create failed in db")
	}
	ErrDBBatchCreate = func(err error) error {
		return errors.Wrap(err, fmt.Sprintf("batch create failed in db"))
	}
	ErrDBSave = func(err error) error {
		return errors.Wrap(err, "save failed in db")
	}
	ErrDBDelete = func(err error) error {
		return errors.Wrap(err, "delete failed in db")
	}
	ErrDBNotOneRowsAffected = errors.New("delete failed, not one rows affected in db")
	ErrDBRecordNotFound     = errors.New("record not found")
	ErrDBFirst              = func(err error) error {
		return errors.Wrap(err, "find failed first in db")
	}
	ErrDBFind = func(err error) error {
		return errors.Wrap(err, "find failed in db")
	}
	ErrDBCount = func(err error) error {
		return errors.Wrap(err, "count failed in db")
	}
	ErrDBUpdate = func(err error) error {
		return errors.Wrap(err, "update failed in db")
	}
	ErrRequestEmptyBody = func() error {
		return errors.New("request body is empty")
	}
	ErrUserIDNotFound = func() error {
		return errors.New("userID not found")
	}
	ErrReadAll = func(err error) error {
		return errors.Wrap(err, "readAll failed")
	}
	ErrParseUint = func(err error) error {
		return errors.Wrap(err, "parseUint failed")
	}
	ErrJSONUnmarshal = func(err error) error {
		return errors.Wrap(err, "json unmarshal failed")
	}
	ErrJSONMarshal = func(err error) error {
		return errors.Wrap(err, "json marshal failed")
	}
	ErrURLNotFound = func() error {
		return errors.New("URL not found")
	}
	ErrURLParse = func(err error) error {
		return errors.Wrap(err, "Parse failed in url service")
	}
)
