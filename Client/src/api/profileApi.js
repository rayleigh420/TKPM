import axios from "../utils/axios"

export const updateProfile = async ({ user_id, name, email, avatar }) => {
    const result = await axios.put(`/user/${user_id}`, {
        name: name,
        avatar: avatar,
        email: email
    })
    return result.data
}