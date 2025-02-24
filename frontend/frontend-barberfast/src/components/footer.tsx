export default function Footer() {
    return (
        <section id="contato" className="bg-mainColor py-[5%] text-white">
            <div className="footer grid grid-cols-1 md:grid-cols-3 gap-8 px-[5%]">
                
                {/* Logo */}
                <div className="box flex flex-col justify-center items-center text-center">
                    <h1 className="font-archivo-black text-[30px]">BARBERFAST</h1>
                </div>

                {/* Contato */}
                <div className="box flex flex-col justify-center items-center text-center">
                    <div className="sub flex items-center gap-2 mt-2">
                        <img src="/icons/footer/location.svg" alt="location" width="25px"/>
                        <p>Rua BarberFast, 123</p>
                    </div>
                    <div className="sub flex items-center gap-2">
                        <img src="/icons/footer/phone.svg" alt="phone" width="25px"/>
                        <p>(11) 99999-9999</p>    
                    </div>
                    <div className="sub flex items-center gap-2 mt-2">
                        <img src="/icons/footer/mail.svg" alt="mail" width="25px"/>
                        <p>contato@barberfast.com</p>
                    </div>
                </div>

                {/* Horário de Funcionamento */}
                <div className="box flex flex-col justify-center items-center text-center">
                    <h1 className="font-archivo-black text-[30px]">ABERTURA</h1>
                    <p>Seg - Sex: 9h - 17h</p>
                    <p>Sáb: 8h - 14h</p>
                </div>

                {/* Suporte */}
                <div className="box flex flex-col justify-center items-center text-center">
                    <h1 className="font-archivo-black text-[30px]">SUPORTE</h1>
                    <p className="cursor-pointer hover:underline">Política de Privacidade</p>
                    <p className="cursor-pointer hover:underline">Termos e Condições</p>
                </div>

                {/* Redes Sociais */}
                <div className="box flex flex-col justify-center items-center text-center">
                    <h1 className="font-archivo-black text-[30px]">SOCIAL</h1>
                    <div className="social flex gap-4 mt-2">
                        <img src="/icons/footer/whatsapp.svg" alt="whatsapp" width="30px" className="cursor-pointer hover:scale-110 transition"/>
                        <img src="/icons/footer/facebook.svg" alt="facebook" width="30px" className="cursor-pointer hover:scale-110 transition"/>
                        <img src="/icons/footer/instagram.svg" alt="instagram" width="30px" className="cursor-pointer hover:scale-110 transition"/>
                    </div>
                </div>

            </div>
        </section>
    );
}
