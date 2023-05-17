import axios from "../utils/axios"

export const getAllBook = async () => {
    const result = await axios.get("/admin/books")
    return result.data
}

export const getGenericBook = async (generic, page) => {
    const result = await axios.get(`/books/${generic}?page=${page}`)
    // console.log(result)
    return {
        books: result.data,
        maxPage: result.headers.total
    }
}

export const getBookType = async (generic, page) => {
    const result = await axios.get(`/books/type?book_type=${generic}&page=${page}}`)
    // console.log(result)
    return {
        books: result.data,
        maxPage: result.headers.total
    }
}

export const getDetailBook = async (id) => {
    const result = await axios.get(`/books/${id}`)
    return result.data
}

export const getVersionBook = async (id) => {
    const result = await axios.get(`/books/${id}/detail`)
    return result.data
}

export const rentBook = async (info) => {
    const result = await axios.post('/rent', info)
    return result.data
}

export const addBook = async (data) => {
    const result = await axios.post('/books', data)
    return result.data
}

export const updateBook = async (data) => {
    const result = await axios.put(`/books/${data.book_id}`, data.info)
    return result.data
}

export const deleteBook = async (book_id) => {
    const result = await axios.delete(`/books/${book_id}`)
    return result.data
}

export const searchBook = async (search_text, page) => {
    const result = await axios.post(`/books/search?page=${page}`, {
        search_text: search_text
    })
    console.log(result)
    return {
        books: result.data,
        maxPage: result.headers.total
    }
}