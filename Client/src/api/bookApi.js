import axios from "../utils/axios"

export const getGenericBook = async (generic, page) => {
    const result = await axios.get(`/books/${generic}?page=${page}`)
    console.log(result)
    return {
        books: result.data,
        maxPage: result.headers.total
    }
}

export const getBookType = async (generic, page) => {
    const result = await axios.get(`/books/type?book_type=${generic}&page=${page}}`)
    console.log(result)
    return {
        books: result.data,
        maxPage: result.headers.total
    }
}