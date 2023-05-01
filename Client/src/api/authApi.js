import axios from "../utils/axios"

export const login = async ({ email, password }) => {
    try {
        const result = await axios.post(`/login`, {
            email: email,
            password: password
        })
        return result.data
    }
    catch (e) {
        console.log(e)
    }
}