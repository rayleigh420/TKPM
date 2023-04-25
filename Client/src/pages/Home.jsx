import BookList from "../components/BookList"

export const data = [

    {
        src: 'https://cdn.wuxiaworld.com/images/covers/hoi.jpg?ver=6ebf4c30b75746d31216c5364893084316be9faa',
        name: 'Hunter of Immortals'
    },
    {
        src: 'https://cdn.wuxiaworld.com/images/covers/bfbt.jpg?ver=fbc27beb0a7017f23af5a9560099d3aeaa2b8d2b',
        name: 'I Became the 1st Floor Boss of the Tower'
    },
    {
        src: 'https://cdn.wuxiaworld.com/images/covers/og.jpg?ver=0104b315d716a06ef770ac6a75c7f00c8fde2918',
        name: 'Overgeared'
    },
    {
        src: 'https://cdn.wuxiaworld.com/images/covers/tnc.jpg?ver=c0680298ea85209c0849028819aeed49338dbead',
        name: "The Nebula's Civilization"
    },
    {
        src: 'https://cdn.wuxiaworld.com/images/covers/fpr.jpg?ver=5b17aa9dc43f6baf7ed45215a1bef4353ac95cd9',
        name: 'The Frozen Player Returns'
    },
    {
        src: 'https://cdn.wuxiaworld.com/images/covers/ar.jpg?ver=b97c8f3393471b676dfc0a360179b41c6c11443b',
        name: 'Absolute Resonance'
    },
    {
        src: 'https://cdn.wuxiaworld.com/images/covers/dr.jpg?ver=8c9af357b85339cbe8bc8b7ad08ee564ee6cf168',
        name: 'Damn Reincarnation'
    }
]

const Home = () => {
    return (
        <>
            <div className="px-[135px] mt-[30px]">
                <BookList nameList="Popular" data={data} />
            </div >
        </>
    )
}

export default Home