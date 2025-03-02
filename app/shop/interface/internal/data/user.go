package data

import (
	"context"

	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/beer-shop/app/shop/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (rp *userRepo) VerifyPassword(ctx context.Context, u *biz.User, password string) error {
	_, err := rp.data.uc.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Username: u.Username,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (rp *userRepo) Find(ctx context.Context, id int64) (*biz.User, error) {
	user, err := rp.data.uc.GetUser(ctx, &usV1.GetUserReq{
		Id: id,
	})
	if err != nil {
		return nil, biz.ErrUserNotFound
	}
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (rp *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	user, err := rp.data.uc.GetUserByUsername(ctx, &usV1.GetUserByUsernameReq{
		Username: username,
	})
	if err != nil {
		return nil, biz.ErrUserNotFound
	}
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (rp *userRepo) Save(ctx context.Context, u *biz.User) error {
	_, err := rp.data.uc.Save(ctx, &usV1.SaveUserReq{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	})
	return err
}

func (rp *userRepo) CreateAddress(ctx context.Context, uid int64, a *biz.Address) (*biz.Address, error) {
	reply, err := rp.data.uc.CreateAddress(ctx, &usV1.CreateAddressReq{
		Uid:      uid,
		Name:     a.Name,
		Mobile:   a.Mobile,
		Address:  a.Address,
		PostCode: a.PostCode,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Address{
		Id:       reply.Id,
		Name:     reply.Name,
		Mobile:   reply.Mobile,
		Address:  reply.Address,
		PostCode: reply.PostCode,
	}, err
}

func (rp *userRepo) GetAddress(ctx context.Context, id int64) (*biz.Address, error) {
	reply, err := rp.data.uc.GetAddress(ctx, &usV1.GetAddressReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Address{
		Id:       reply.Id,
		Name:     reply.Name,
		Mobile:   reply.Mobile,
		Address:  reply.Address,
		PostCode: reply.PostCode,
	}, nil
}

func (rp *userRepo) ListAddress(ctx context.Context, uid int64) ([]*biz.Address, error) {
	reply, err := rp.data.uc.ListAddress(ctx, &usV1.ListAddressReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Address, 0)
	for _, x := range reply.Results {
		rv = append(rv, &biz.Address{
			Id:       x.Id,
			Name:     x.Name,
			Mobile:   x.Mobile,
			Address:  x.Address,
			PostCode: x.PostCode,
		})
	}
	return rv, nil
}

func (rp *userRepo) CreateCard(ctx context.Context, uid int64, c *biz.Card) (*biz.Card, error) {
	reply, err := rp.data.uc.CreateCard(ctx, &usV1.CreateCardReq{
		Uid:     uid,
		CardNo:  c.CardNo,
		Ccv:     c.CCV,
		Expires: c.Expires,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Card{
		Id: reply.Id,
	}, nil
}

func (rp *userRepo) GetCard(ctx context.Context, id int64) (*biz.Card, error) {
	reply, err := rp.data.uc.GetCard(ctx, &usV1.GetCardReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Card{
		Id:     reply.Id,
		CardNo: reply.CardNo,
	}, nil
}

func (rp *userRepo) ListCard(ctx context.Context, uid int64) ([]*biz.Card, error) {
	reply, err := rp.data.uc.ListCard(ctx, &usV1.ListCardReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}

	rv := make([]*biz.Card, 0)
	for _, x := range reply.Results {
		rv = append(rv, &biz.Card{
			Id:     x.Id,
			CardNo: x.CardNo,
		})
	}
	return rv, nil
}
