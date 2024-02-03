<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Str;

class LinkSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        DB::table('link')->insert([
            'nome' => Str::random(10),
            'codigo_hash' => Hash::make('teste'),
            'password' => Hash::make('password'),
            'link_whatsapp' => Str::random(10),
            'link_telegram' => Str::random(10),
            'ordem_de_redirecionamento' => 'telegram,whatsapp',
            'usuario' => 1,
        ]);
    }
}
