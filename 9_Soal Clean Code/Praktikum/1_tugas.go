// 1. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
//     - Berapa banyak kekurangan dalam penulisan kode tersebut? => menurut saya ada 6
//     - Bagian mana saja terjadi kekurangan tersebut? sudah saya tandai di komentar bawah
//     - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.

package main

//1. Penulisan nama struct seharusnya disarankan menggunakan huruf besar jadinya User agar bisa membedakan mana struct mana bukan
type user struct {
	id       int
	username int
	password int
}

//2. Penulisannya sebaiknya menggunakan camelCase untuk nama type struct nya karena sesuai dengan kebiasaan programmer
// Contoh : userService
type userservice struct {
	//Penamaan variabel tipe variabelnya tidak jelas maksudnya apa
	t []user
}

//3. Penulisan methd getallusers() sebaiknya menggunakan camelCase karena sesuai dengan kebiasaan programmer
//contoh yang disarankan : getAllUsers()

//4. pemberian variabel u juga tidak jelas maksudnya apa
func (u userservice) getallusers() []user {
	return u.t
}

//5. Penulisan methd getuserbyid() sebaiknya menggunakan camelCase karena sesuai dengan kebiasaan programmer
//contoh yang disarankan : getUserById()
//6. pemberian variabel u & r juga tidak jelas maksudnya apa
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}