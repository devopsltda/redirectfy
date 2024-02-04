<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::connection('sqlite')->create('usuario', function (Blueprint $table) {
            $table->id()->unique();
            $table->string('cpf')->unique();
            $table->string('nome');
            $table->string('nome_de_usuario')->unique();
            $table->string('email')->unique();
            $table->string('senha');
            $table->dateTimeTz('data_de_nascimento');
            $table->foreignId('plano_de_assinatura');
            $table->timestampsTz();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('usuario');
    }
};
