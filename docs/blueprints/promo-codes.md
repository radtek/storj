# Promo Codes 

## Abstract

Enable users to apply a promotional code to their account so that they can automatically receive the free credit associated with that code.

Summary: When Sales/Marketing whats to run a promotion or hand out a specific amount of free credit to a group of people, we want to be able to generate a promo code that a user could enter into their account, so that specific free credits and expiration of the credits would be automatically applied to that user’s account.

## Background


## Design

There are 2 high level aspects of this design:

### Generating Promo Codes

The satellite admin can generate promo codes from the admin UI. Once generated, these promo codes can be applied to any user's account.

### Applying a Promo Code to a User

On account creation or from the payments panel, a user should be able to insert a promo code and immediately have their project usage limits increased. If the promo code has an expiration associated with it, these limit increases and this promo code will disappear after the expiration date.

## Rationale

In the design detailed below, two additional database tables are defined. In theory, we could just have the `promo_codes` table and add some additional fields to the `users` table to get rid of the `user_promo_codes` table requirement. However, this will put more load on the `users` table, and deleting promo codes from users will have worse performance.

## Implementation

### Database Layer

We will need a new table for promo codes, called `promo_codes`. The table will have the following fields:
```
name        string (unique)
expireAfter uint (hours or days or something) (nullable)
bandwidth   uint64
storage     uint64
```
`name` must be a unique field for this table. It can be used as the key for the table.

We will need another table to indicate which users have which promo codes associated with them. We will call this `user_promo_codes`. This will have the following fields:
```
promo_code_name string
user_id         id
expiration      time (nullable)
bandwidth       uint64
storage         uint64
```
`user_promo_codes` must have a constraint allowing only one unique `user_id` at a time. A single user can only have one promo code associated with their account. `promo_code_name` references the `name` column in the `promo_codes` table. There must be a row with the same name in the `promo_codes` table.

### Service

The service layer will communicate with the database layer to add, edit, and remove promo codes as an admin, or apply a promo code as a user.

Here are some type definitions for the service layer:
```
type PromoCode struct {
    name        string
    expiration  *time.Duration
    bandwidth   memory.Size
    storage     memory.Size
}
type UserPromoCode struct {
    name        string
    expiration  *time.Time
    bandwidth   memory.Size
    storage     memory.Size
}
```

It will require the following interface:

```
AddPromoCode(newPromoCode PromoCode) error
UpdatePromoCode(updatedPromoCode PromoCode) error
DeletePromoCode(promoCodeName string) error

ApplyPromoCodeToUser(userID uuid.UUID, promoCodeName string, now time.Time) (newBandwidthLimit, newStorageLimit memory.Size, err error)
GetUserPromoCode(userID uuid.UUID) (*UserPromoCode, error)
RemoveExpiredPromoCodesFromUsers(now time.Time) error
```

* `AddPromoCode` - insert a new row into the `promo_codes` table with the provided values. Return an error if a promo code with the provided `name` already exists.
* `UpdatePromoCode` - update a rew in the `promo_codes` table with the provided values. Return an error if a promo code with the provided `name` does not exist.
* `DeletePromoCode` - first, for each user using the promo code provided, subtract the bandwidth/storage values for that promo code from the user's project limits. Then, delete the promo code's row from the `promo_codes` table.
* `ApplyPromoCodeToUser` - First, add a row to the `user_promo_codes` table with the provided user and promo code. Return an error if either doesn't exist in the database. Then, increase the user's project bandwidth/storage limits according to the promo code's information. If `promoCode.Expiration` is not null, set the `user_promo_codes` expiration as `now+promoCode.Expiration`.
* `GetUserPromoCode` - If it exists, return the information for a user from the `user_promo_codes` table. Otherwise, return `nil`.
* `RemoveExpiredPromoCodesFromUsers` - For each `user_promo_codes` row with a non-null expiration time before `now`, decrease bandwidth/storage according to the promo code limit increases defined in `user_promo_codes`. Then, delete the row from `user_promo_codes`.

**Important note** - Any time a user's limits are decreased as a result of expiration or promo code deletion, it is very important that the values used to decrease limits are coming from `user_promo_codes`, and not `promo_codes` - a promo code may be updated to change its limits, but limits will not be updated for users who have already applied the promo code.

### Chore

A chore will be necessary to remove expired promo codes from users and update their project limits. It will run on a configured interval `promocodes.ExpirationChoreInterval`, which can be set to once a day, once every few hours, or whatever we want aver discussing it more. It shouldn't be very performance intensive. All it will do is call `service.RemoveExpiredPromoCodesFromUsers()`.

### UI
#### **Admin**
The satellite admin should be able to add new promo codes from the UI. There should be fields for name, expiration, and bandwidth/storage limit increases. Expiration should be an optional field.

It should be possible to delete existing promo codes from the UI.

Editing existing promo codes is not necessary for a minimal implementation, but in the future, we may want the ability to edit promo codes from the UI.

#### **User**
There are a couple options we have from a UI perspective here. Either a user can create a promo code on account creation, or add it from the payments panel in the webapp. We may even end up having both enabled or AB test one vs. the other. Either way, all that needs to be done is make a call to `service.ApplyPromoCodeToUser` when a user attempts to apply a promo code.

## Wrapup

The Satellite Web Team will take ownership of this project and is responsible for archiving this blueprint upon completion.

## Open issues

* Can a user delete a promo code from their account?
* Difference between project vs. user limits? Design may need to be updated after more research.
* Is editing promo codes necessary? Should it be removed from the design?
* Is there existing documentation that needs to be updated? (add to wrapup)
* Is there new documentation that needs to be written? (add to wrapup)