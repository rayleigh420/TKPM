import { data } from "../pages/home"
import BookList from "./list/BookList"

const RelatedBook = () => {
    return (
        <div className="flex flex-col items-start">
            <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff]'>Related Book</p>
            <BookList data={data} />
        </div>
    )
}

export default RelatedBook
