package errors

import "errors"

var ErrInvalidRequest = errors.New("invalid request")
var ErrLoginFailed = errors.New("wrong email or password")
var ErrEmailAlready = errors.New("email already exist")
var ErrReffCodeInvalid = errors.New("refferal code invalid")

var ErrUserNotFound = errors.New("user not found")
var ErrSizeNotFound = errors.New("size not found")
var ErrCategoryNotFound = errors.New("category not found")
var ErrAddOnNotFound = errors.New("add on not found")
var ErrAddressNotFound = errors.New("address not found")
var ErrShippingNotFound = errors.New("shipping not found")
var ErrPaymentNotFound = errors.New("payment not found")
var ErrPromoNotFound = errors.New("promo not found")
var ErrPromoUserNotFound = errors.New("promo user not found")

var ErrMinTopUp = errors.New("topup min Rp.10.000")
var ErrMaxTopUp = errors.New("topup max Rp.10.000.000")

var ErrMinAddOns = errors.New("min select 1 add on")
var ErrAddOnInvalid = errors.New("add on id invalid")

var ErrShippingMustPaid = errors.New("shipping must be paid first")
var ErrShippingStatus = errors.New("status only: PROCESS, PICKUP, DELIVERY, DELIVERED")
var ErrShippingReview = errors.New("send review after shipping done")
var ErrShippingAlreadyPaid = errors.New("shipping has been paid")
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrPromoAlreadyUsed = errors.New("promo already used")
var ErrPromoFeeInvalid = errors.New("promo promo fee invalid")
var ErrQuotaOutOfStock = errors.New("quota out of stock")
var ErrQuotaExpired = errors.New("quota expired")

var ErrGameTransactionNotDone = errors.New("shipping and payment must be done first")
var ErrGameMinTransaction = errors.New("transaction min Rp20.000")
var ErrGameChanceUsed = errors.New("chance play game has been used")
