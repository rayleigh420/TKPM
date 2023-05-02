import imgAxios from "../utils/imgAxios";

export const uploadAva = async (imgForm) => {

    const result = await imgAxios.post('', imgForm);
    return result.data.data.url
}