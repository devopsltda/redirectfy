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
        Schema::connection('sqlite')->create('link', function (Blueprint $table) {
            $table->id();
            $table->string('nome');
            $table->string('codigo_hash');
            $table->string('link_whatsapp');
            $table->string('link_telegram');
            $table->string('ordem_de_redirecionamento');
            $table->foreignId('usuario');
            $table->timestampsTz();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('link');
    }
};
