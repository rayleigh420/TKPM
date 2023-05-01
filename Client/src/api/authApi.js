import axios from "../utils/axios"

export const login = async ({ email, password }) => {
    // try {
    // }
    // catch (e) {
    //     throw Error(e)
    // }
    const result = await axios.post(`/login`, {
        email: email,
        password: password
    })
    return result.data
}

export const signUp = async ({ name, phone, email, password }) => {
    // try {
    // }
    // catch (e) {
    //     console.log(e)
    // }
    console.log(name, phone, email, password)
    const result = await axios.post(`/signup`, {
        name: name,
        phone: phone,
        email: email,
        password: password
    })
    return result.data
}