import axios from "../utils/axios"

export const getGenericBook = async (generic, page) => {
    const result = await axios.get(`/books/${generic}?page=${page}`)
    console.log(result)
    return result.data
}