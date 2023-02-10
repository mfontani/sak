#!/usr/bin/env perl
use 5.020_000;
use warnings;
use Mojo::UserAgent qw<>;
use autodie;

my $page = '';
{
    my $res = Mojo::UserAgent->new->get('https://www.nerdfonts.com/cheat-sheet')->result;
    if ($res->is_success) {
        $page = $res->body;
    }
    elsif ($res->is_error) {
        die $res->message;
    }
    else {
        die "AIEE";
    }
}
# Example block:
##  <div class="column">
##    <div class="nf nf-cod-bug center"></div>
##    <div class="class-name">nf-cod-bug</div><div class="codepoint">eaaf</div>
##  </div>
my $RX_BLOCK = qr{
    <div \s+ class="column">
        \s*
        (?:
            <span \s class="corner-red"></span><span \s class="corner-text">(?<obsolete>obsolete)</span>
            \s*
        )?
        <div \s+ class="[^"]+"></div>
        \s*
        <div \s+ class="class-name">(?<name>[^<]+)</div>
        \s*
        <div \s+ title="[^"]+" \s+ class="codepoint">(?<hex>[1-9a-fA-F][0-9a-fA-F]*)</div>
        \s*
    </div>
}xms;
open my $fh, '>', 'nf.go';
printf {$fh} "package main\n\n";
printf {$fh} "var nfRunes = map[rune]string{\n";
my %found;
while ($page =~ m!$RX_BLOCK!xmsg) {
    my ($hex, $name, $obsolete) = ($+{hex}, $+{name}, $+{obsolete});
    $obsolete = defined $obsolete && length $obsolete ? " ($obsolete)" : '';
    $found{$hex}++ and next;
    printf {$fh} sprintf qq!\t%-8s"%s%s",\n!,
        (sprintf '%d:', hex("0x$hex")),
        $name, $obsolete;
}
printf {$fh} "}\n";
close $fh;
