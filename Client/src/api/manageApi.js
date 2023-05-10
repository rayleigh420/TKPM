import axios from "../utils/axios"

export const getBookedBook = async () => {
    const result = await axios.get('/rentlist')
    return result.data
}

export const searchBookedBook = async (bookedID) => {
    console.log(bookedID)
    const result = await axios.get(`/rentlist/search?search_id=${bookedID}`)
    return result.data
}

export const getBorrowedBook = async () => {
    const result = await axios.get('/borrowlist')
    return result.data
}

export const getPaidBook = async () => {
    const result = await axios.get('/returnlist')
    return result.data
}