// import imgAxios from "../utils/imgAxios";
import axios from "axios";

export const uploadAva = async (imgForm) => {
    console.log(imgForm.get("image"))

    const result = await axios.post('https://api.imgbb.com/1/upload?key=YOUR_CLIENT_API_KEY', imgForm);

    console.log(result)
    return result.data.data.url
}

// import { imgbbUploader } from "imgbb-uploader"


// export const uploadAva = async (imgUrl) => {
//     const options = {
//         apiKey: "API_KEY", // MANDATORY
//         imageUrl: imgUrl, // OPTIONAL: pass an URL to imgBB (max 32Mb)
//     };
//     console.log(imgUrl)
//     const result = await imgbbUploader(options)
//     console.log(result)
//     return result
// }
