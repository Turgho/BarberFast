'use client';

import { Swiper, SwiperSlide } from 'swiper/react';
import { Navigation, Pagination, Autoplay } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';

const slides = [
    { bgClass: 'bg-home-1', text: 'Bem-vindo à', span: 'BARBERFAST' },
    { bgClass: 'bg-home-2', text: 'cortes', span: 'MELHORES' },
    { bgClass: 'bg-home-3', text: 'Promoções', span: 'IMPERDÍVEIS' },
];

export default function Carrosel(){
    return (
        <Swiper
            modules={[Navigation, Pagination, Autoplay]}
            spaceBetween={0}
            slidesPerView={1}
            pagination={{ clickable: true }}
            autoplay={{ delay: 4000 }}
            loop
            className="w-[100%] h-[600px]"
        >
        {slides.map((slide, index) => (
            <SwiperSlide key={index}>
                {/* Imagem de fundo */}
                <div className={`flex justify-center items-center w-full h-full ${slide.bgClass} bg-cover bg-center relative`}>
                    {/* Overlay com texto */}
                    <div className="overlay flex flex-col md:flex-row justify-center items-center w-full h-full bg-mainColor/50 text-base md:text-[30px] text-center font-archivo-black absolute flex-wrap leading-relaxed">
                        {
                            index === 1 ? (
                                <>
                                    <span className="text-laranja">{slide.span}</span>&nbsp;{slide.text}!
                                </>
                            ) : (
                                <>
                                    {slide.text}&nbsp;<span className="text-laranja">{slide.span}</span>!
                                </>
                            )
                        }
                    </div>
                </div>
            </SwiperSlide>
        ))}
        </Swiper>
    );
}
