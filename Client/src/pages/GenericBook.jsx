import { useParams } from "react-router-dom"
import GenericBookList from "../components/list/GenericBookList"

const GenericBook = () => {
    const { generic } = useParams()
    return (
        <div className="px-[216px] mt-[30px]">
            <GenericBookList />
        </div>
    )
}

export default GenericBook