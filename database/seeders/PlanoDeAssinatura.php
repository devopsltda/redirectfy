<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;

class PlanoDeAssinatura extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        DB::table('plano_de_assinatura')->insert([
            'nome' => Str::random(10),
            'valor_mensal' => random_int(10000, 100000),
        ]);
    }
}
