import axios from "../utils/axios"

export const getPaidBook = async () => {
    const result = await axios.get('/history')
    return result.data
}