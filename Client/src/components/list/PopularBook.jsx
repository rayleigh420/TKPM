import { useQuery } from "@tanstack/react-query"
import { data } from "../../pages/home"
import BookList from "./BookList"
import { getGenericBook } from "../../api/bookApi"

const PopularBook = () => {
    const { data, isLoading, isError } = useQuery({
        queryKey: ["books", "popular", 1],
        queryFn: () => getGenericBook("popular", 1)
    })

    // console.log(data)
    return (
        <div className="flex flex-col items-start mt-[50px]">
            <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff]'>Populars</p>
            <BookList data={data.slice(0, 7)} />
        </div>
    )
}

export default PopularBook