import axios from "../utils/axios"

export const getGenericBook = async (generic, page = 1) => {
    const result = await axios.get(`/books/${generic}?page=${page}`)
    return result.data
}