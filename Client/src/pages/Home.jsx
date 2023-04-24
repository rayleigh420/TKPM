const data = [

    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    },
    {
        src: 'https://cdn.wuxiaworld.com/images/covers/bfbt.jpg?ver=fbc27beb0a7017f23af5a9560099d3aeaa2b8d2b',
        name: 'I Became the 1st Floor Boss of the Tower'
    },
    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    },
    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    },
    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    },
    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    },
    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    },
    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    },
    {
        src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
        name: 'Boi hoi gia'
    }
]

const Home = () => {
    return (
        <>
            <div className="px-[135px] mt-[30px]">
                <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff]'>Popular</p>

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
            </div >
        </>
    )
}

export default Home