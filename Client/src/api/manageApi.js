import axios from "../utils/axios"

export const getBorrowedBook = async () => {
    const result = await axios.get('/borrowlist')
    return result.data
}

export const getPaidBook = async () => {
    const result = await axios.get('/returnlist')
    return result.data
}