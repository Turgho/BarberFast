'use client';

import Link from "next/link";
import Button from "@/components/button";
import Carrosel from "@/components/carrosel";

export default function Index(){
    return (
        <section id="index">
            <div className="flex flex-col md:flex-row sm:flex-wrap min-h-screen w-screen justify-center items-center gap-10 bg-mainColor p-4">
                <div className="box-esquerda flex flex-col gap-[80px] mt-[30%] md:mt-0">
                    <div className="textos flex flex-col gap-4">
                        <h1 className="text-[55px] text-center leading-none">
                            A <span>ARTE</span> DE BEM <br/><span>CUIDAR</span> DO <span className="text-laranja">HOMEM</span>.
                        </h1>

                        <p className="text-center">
                            Faça cadastro para agendar seu horário, ficar por <br /> dentro de promoções e receber avisos!
                        </p>
                    </div>
                    <div className="agendar flex flex-col justify-center items-center gap-2">
                        <Link href={`/agendar`}>
                            <Button text="AGENDAR!" variant="primary" size="lg"/>
                        </Link>
                        <Link href="/login" className="underline underline-offset-2 text-[70%]">Já é cliente? Faça login</Link>
                    </div>
                </div>
 
                <div className="box-direita w-full md:w-[30%] shadow-custom order-last">
                    <Carrosel />
                </div>
            </div>
        </section>
    );
}
