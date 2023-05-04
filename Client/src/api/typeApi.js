import axios from "../utils/axios"

export const getListType = async () => {
    const result = await axios.get('types')
    return result.data
}