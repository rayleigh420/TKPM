import { data } from "../pages/home"
import BookList from "./BookList"

const PopularBook = () => {
    return (
        <div className="flex flex-col items-start mt-[50px]">
            <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff]'>Populars</p>
            <BookList data={data} />
        </div>
    )
}

export default PopularBook