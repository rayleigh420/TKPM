import axios from "../utils/axios"

export const login = async ({ email, password }) => {
    try {
        console.log(email, password)
        const result = await axios.post(`/login`, {
            email: email,
            password: password
        })
        console.log(result)
        return result.data
    }
    catch (e) {
        console.log(e)
    }
}

export const signUp = async ({ name, phone, email, password }) => {
    try {
        console.log(email, password)
        const result = await axios.post(`/signup`, {
            name: name,
            phone: phone,
            email: email,
            password: password
        })
        console.log(result)
        return result.data
    }
    catch (e) {
        console.log(e)
    }
}