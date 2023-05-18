import { Link } from "react-router-dom"
import { data } from "../../pages/home"
import BookList from "./BookList"
import { useQuery } from "@tanstack/react-query"
import { getGenericBook } from "../../api/bookApi"

const NewestBook = () => {

    const { data, isLoading, isError } = useQuery({
        queryKey: ["books", "newest", 1],
        queryFn: () => getGenericBook("newest", 1)
    })

    return (
        <div className="flex flex-col items-start mt-[30px]">
            <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff]'>Newest</p>
            <BookList data={data} isLoading={isLoading} />
            {/* <div className="carousel carousel-center max-w-[370px] h-[430px] p-4 space-x-4 bg-neutral rounded-box">
                {data.map(item => (
                    <div className="carousel-item w-[220px] flex flex-col gap-[20px] cursor-pointer">
                        <Link to='book'>
                            <img src={item.src} className="rounded-box w-full h-[320px]" />
                            <div className="w-full text-center text-[#ffffff] font-semibold">{item.name}</div>
                        </Link>
                    </div>
                ))}
            </div> */}
            {/* <BookList data={data} /> */}
        </div>
    )
}

export default NewestBook 