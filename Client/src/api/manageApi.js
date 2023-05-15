import axios from "../utils/axios"

export const getBookedBook = async () => {
    const result = await axios.get('/rentlist')
    return result.data
}

export const searchBookedBook = async (bookedID) => {
    const result = await axios.get(`/rentlist/search?search_id=${bookedID}`)
    return result.data
}

export const getBorrowedBook = async () => {
    const result = await axios.get('/borrowlist')
    return result.data
}

export const searchBorrowedBook = async (borrowedID) => {
    const result = await axios.get(`/borrowlist/search?search_id=${borrowedID}`)
    console.log(result.data)
    return result.data
}

export const getPaidBook = async () => {
    const result = await axios.get('/returnlist')
    return result.data
}

export const addVersionBook = async (info) => {
    console.log("API:", info)
    const result = await axios.post(`/version/${info.book_id}`, {
        location: info.location
    })
    return result.data
}