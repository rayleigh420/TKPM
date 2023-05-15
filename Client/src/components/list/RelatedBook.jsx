
import BookList from "./BookList"

const RelatedBook = ({ related, isLoading }) => {
    // console.log(related)
    return (
        <div className="flex flex-col items-start mb-[30px]">
            <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff]'>Related Book</p>
            <BookList data={{ books: related }} isLoading={isLoading} />
        </div>
    )
}

export default RelatedBook
