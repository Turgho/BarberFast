'use client';

import Navbar from "@/components/navbar";
import Index from "@/components/index"
import Sobre from "@/components/sobre"
import Servicos from "@/components/servicos";
import Precos from "@/components/precos";
import Avaliacoes from "@/components/avaliacoes";
import Galeria from "@/components/galeria";
import Footer from "@/components/footer";
import Copyright from "@/components/copyright";

import Head from "next/head";
import "@/styles/globals.css"

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
      <html lang="pt-BR">
        <Head>
            <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Julius+Sans+One&display=swap"/>
        </Head>
        <body>
          <Navbar/>
          <Index/>
          <Sobre/>
          <Servicos/>
          <Precos/>
          <Avaliacoes/>
          <Galeria/>
          <Footer/>
          <Copyright/>
          {children}
        </body>
      </html>
    );
  }
  