const BookList = ({ data }) => {
    return (
        <div className="flex gap-5 flex-wrap mt-[20px]">
            {
                data.map(item => (
                    <div className='w-[158px] h-[270px] rounded-[7px] cursor-pointer'>
                        <img src={item.src} className='w-full h-[230px]' />
                        <p className='text-[#e0e0e0] text-[15.2px] font-semibold leaiding-[18.24px] tracking-[-0.304px] mt-1'>{item.name}</p>
                    </div>
                ))
            }

        </div>
    )
}

export default BookList