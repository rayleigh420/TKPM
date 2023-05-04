const BookList = ({ data, isLoading }) => {
    return (
        <div className="flex gap-5 flex-wrap mt-[20px]">
            {
                isLoading ?
                    (
                        [...Array(7)].map(item => (
                            <div className='w-[158px] h-[270px] rounded-[7px] cursor-pointer'>
                                <div className="w-full h-[230px] bg-gray-200">
                                    {/* <svg class="w-12 h-12 text-gray-200" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" fill="currentColor" viewBox="0 0 640 512"><path d="M480 80C480 35.82 515.8 0 560 0C604.2 0 640 35.82 640 80C640 124.2 604.2 160 560 160C515.8 160 480 124.2 480 80zM0 456.1C0 445.6 2.964 435.3 8.551 426.4L225.3 81.01C231.9 70.42 243.5 64 256 64C268.5 64 280.1 70.42 286.8 81.01L412.7 281.7L460.9 202.7C464.1 196.1 472.2 192 480 192C487.8 192 495 196.1 499.1 202.7L631.1 419.1C636.9 428.6 640 439.7 640 450.9C640 484.6 612.6 512 578.9 512H55.91C25.03 512 .0006 486.1 .0006 456.1L0 456.1z" /></svg> */}
                                </div>
                                {/* <img src={item.book_img} className='w-full h-[230px]' /> */}
                                <div class="h-1.5 bg-gray-200 rounded-full dark:bg-gray-700 max-w-[130px] mt-2"></div>
                                <div class="h-1.5 bg-gray-200 rounded-full dark:bg-gray-700 max-w-[145px] mt-1"></div>
                            </div>
                        ))
                    )
                    :
                    (
                        data.books && data.books.slice(0, 7).map(item => (
                            <div className='w-[158px] h-[270px] rounded-[7px] cursor-pointer'>
                                <img src={item.book_img} className='w-full h-[230px]' />
                                <p className='text-[#e0e0e0] text-[15.2px] font-semibold leaiding-[18.24px] tracking-[-0.304px] mt-1'>{item.name}</p>
                            </div>
                        ))
                    )
            }
            {/* {
                data && data.map(item => (
                    <div className='w-[158px] h-[270px] rounded-[7px] cursor-pointer'>
                        <img src={item.book_img} className='w-full h-[230px]' />
                        <p className='text-[#e0e0e0] text-[15.2px] font-semibold leaiding-[18.24px] tracking-[-0.304px] mt-1'>{item.name}</p>
                    </div>
                ))
            } */}

        </div>
    )
}

export default BookList