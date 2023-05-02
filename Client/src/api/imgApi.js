import imgAxios from "../utils/imgAxios";

export const uploadAva = async (img) => {

    const result = await imgAxios.post('', img);
    console.log(result)
    return result.data.data.url
}