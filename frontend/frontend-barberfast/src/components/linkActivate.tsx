import { useState, useEffect } from 'react';
import Link from 'next/link';

interface LinkProps {
  href: string;
  name: string;
  activeSection: string;  // Adicionando a prop activeSection
}

export default function ActivateLink({ href, name, activeSection }: LinkProps) {
  const isActive = activeSection === href.substring(2); // Remove o "#" para comparar com o ID da seção

  return (
      <Link 
          href={href} 
          className={`duration-300 animate-fadeIn ${isActive ? 'text-laranja font-bold' : 'text-white hover:text-laranja'}`}
      >
          {name}
      </Link>
  );
}