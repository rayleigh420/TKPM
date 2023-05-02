import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { Link, useParams } from "react-router-dom"
import { getGenericBook } from "../../api/bookApi"

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]

const GenericBookList = () => {
    const { generic } = useParams();
    const [page, setPage] = useState(1)

    const { data: books, isLoading, isFetching, isError } = useQuery({
        queryKey: ['books', page],
        queryFn: () => getGenericBook(generic, page)
    })

    // console.log(books[0].bookItems)

    return (
        <div className="flex flex-col items-start mt-[50px]">
            <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff] capitalize'>{generic}</p>
            <div className="w-full mt-[20px] grid grid-cols-2 gap-y-6 gap-x-3">
                {
                    books && books[0].bookItems.map(item => (
                        <div className="flex flex-row gap-[12px]" key={item.book_id}>
                            <Link to="/book">
                                <div className={`bg-[url('${item.book_img}')]` + " cursor-pointer min-w-[125px] h-[180px] bg-cover hover:border-[0.1px] hover:border-[#142B45] rounded-[7px]"}></div>
                            </Link>
                            <div className="">
                                <div className="badge bg-[#eeeeee] rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Ongoing</div>
                                <Link to="/book">
                                    <h1 className="mt-[5px] text-[#ffffff] font-bold leading-[20px] cursor-pointer">{item.name}</h1>
                                </Link>
                                <p className="mt-[10px] text-[#bdbdbd] text-[14px] font-medium">Author: {item.author}</p>
                                <p className="text-ellipsis overflow-hidden mt-[10px] text-[#bdbdbd] text-[14px] font-medium">
                                    For decades, he had lived as a puppet of the tower. Throughout that time, who knew how many deaths he had encountered? One day, however, his memories returned to him. The First Floor Boss finally awakened
                                </p>
                            </div>
                        </div>
                    ))
                }
            </div>
            <div className="flex flex-row justify-center mt-[30px] w-full">
                <div className="btn-group">
                    <button className="btn">1</button>
                    <button className="btn btn-active">2</button>
                    <button className="btn">3</button>
                    <button className="btn">4</button>
                </div>
            </div>
        </div>
    )
}

export default GenericBookList