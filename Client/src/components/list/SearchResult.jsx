import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { Link, useLocation, useParams } from "react-router-dom"
import { searchBook } from "../../api/bookApi";

const SearchResult = () => {
    const [page, setPage] = useState(1)

    const location = useLocation()
    const searchParams = new URLSearchParams(location.search);
    const encodedKeyword = searchParams.get('keyword');
    const keyword = decodeURIComponent(encodedKeyword);

    console.log(keyword)

    const queryClient = useQueryClient()

    const { data, isLoading, isFetching, isError } = useQuery({
        queryKey: ['search', keyword, page],
        queryFn: () => searchBook(keyword, page),
        onSuccess: () => {
            window.scrollTo({ top: 0, behavior: 'smooth' });
        },
        keepPreviousData: true
    })

    const prefetchPrevPage = (page) => {
        queryClient.prefetchQuery(['books', keyword, page], {
            queryFn: () => searchBook(keyword, page),
            staleTime: 10 * 1000,
            cacheTime: 10 * 1000,
        })
    }

    const prefetchNextPage = (page) => {
        queryClient.prefetchQuery(['books', keyword, page], {
            queryFn: () => searchBook(keyword, page),
            staleTime: 10 * 1000,
            cacheTime: 10 * 1000,
        })
    }

    console.log(data)

    return (
        <div className="px-[216px] mt-[30px]">

            <div className="flex flex-col items-start mt-[50px]">
                <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff] capitalize'>Search {keyword}:</p>
                <div className="w-full mt-[20px] grid grid-cols-2 gap-y-6 gap-x-3">
                    {
                        isLoading ?
                            (
                                [...Array(6)].map(item => (
                                    <div className="flex flex-row gap-[12px]">
                                        <div >
                                            <svg className="w-[125px] h-[180px] text-gray-200" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" fill="currentColor" viewBox="0 0 640 512"><path d="M480 80C480 35.82 515.8 0 560 0C604.2 0 640 35.82 640 80C640 124.2 604.2 160 560 160C515.8 160 480 124.2 480 80zM0 456.1C0 445.6 2.964 435.3 8.551 426.4L225.3 81.01C231.9 70.42 243.5 64 256 64C268.5 64 280.1 70.42 286.8 81.01L412.7 281.7L460.9 202.7C464.1 196.1 472.2 192 480 192C487.8 192 495 196.1 499.1 202.7L631.1 419.1C636.9 428.6 640 439.7 640 450.9C640 484.6 612.6 512 578.9 512H55.91C25.03 512 .0006 486.1 .0006 456.1L0 456.1z" /></svg>
                                        </div>
                                        <div className="">
                                            {/* <div className="badge bg-[#eeeeee] rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Ongoing</div> */}
                                            <div className="h-2.5 bg-gray-200 rounded-full dark:bg-gray-700 w-48 mb-4 mt-[35px]"></div>
                                            <div className="h-2 bg-gray-200 rounded-full dark:bg-gray-700 max-w-[480px] mb-2.5"></div>
                                            <div className="h-2 bg-gray-200 rounded-full dark:bg-gray-700 mb-2.5"></div>
                                            <div className="h-2 bg-gray-200 rounded-full dark:bg-gray-700 w-[300px] mb-2.5"></div>
                                            <div className="h-2 bg-gray-200 rounded-full dark:bg-gray-700 w-[320px] mb-2.5"></div>
                                            <div className="h-2 bg-gray-200 rounded-full dark:bg-gray-700 w-[360px]"></div>
                                        </div>
                                    </div >
                                ))
                            )
                            :
                            (
                                data.books && data.books?.map(item => (
                                    <div className="flex flex-row gap-[12px]" key={item.book_id}>
                                        <Link to={"/book/" + item.book_id}>
                                            <div className=" cursor-pointer min-w-[125px] max-w-[130px] h-[180px] bg-cover hover:border-[0.1px] hover:border-[#142B45] rounded-[7px] overflow-hidden">
                                                <img src={item?.book_img || item?.book_image} className="w-full h-full" />
                                            </div>
                                            {/* <div className={`bg-[url('${item.book_img}')]` + " cursor-pointer min-w-[125px] h-[180px] bg-cover hover:border-[0.1px] hover:border-[#142B45] rounded-[7px]"}></div> */}
                                        </Link>
                                        <div className="max-h-[180px] text-ellipsis overflow-hidden">
                                            {
                                                item?.amount > 0
                                                    ?
                                                    <div className="badge bg-[#eeeeee] rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Ongoing</div>
                                                    :
                                                    <div className="badge bg-red-500 rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Over</div>

                                            }
                                            <Link to={"/book/" + item.book_id}>
                                                <h1 className="mt-[5px] text-[#ffffff] font-bold leading-[20px] cursor-pointer">{item.name}</h1>
                                            </Link>
                                            <p className="mt-[15px] text-[#bdbdbd] text-[14px] font-medium">Author: {item.author}</p>
                                            <p className="text-ellipsis overflow-hidden mt-[10px] pr-[5px] text-[#bdbdbd] text-[14px] font-medium">
                                                {item.description}
                                            </p>
                                        </div>
                                    </div>
                                ))
                            )

                    }
                </div >
                <div div className="flex flex-row justify-center mt-[100px] w-full" >
                    <div className="btn-group grid grid-cols-3">
                        <button className="btn hover:bg-[#661ae6] hover:text-[#ffffff] hover:font-bold" onClick={() => setPage(prev => prev - 1)} disabled={page == 1} onMouseEnter={() => prefetchPrevPage(page - 1)}>Previous page</button>
                        <button className="btn bg-[#242933] text-[#ffffff] font-bold hover:bg-[#242933]">Page {page}</button>
                        <button className="btn hover:bg-[#661ae6] hover:text-[#ffffff] hover:font-bold" onClick={() => setPage(prev => prev + 1)} disabled={page == data?.maxPage} onMouseEnter={() => prefetchNextPage(page + 1)}>Next</button>
                    </div>
                </div >
            </div >
        </div>
    )
}

export default SearchResult
