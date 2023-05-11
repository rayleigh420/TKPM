import axios from "../utils/axios"

export const getAllUser = async (token) => {
    const result = await axios.get('/admin/users', {
        token: token
    })
    return result.data
}

export const deleteUser = async (idUser) => {
    const result = await axios.delete(`/user/${idUser}`)
    return result.data
}