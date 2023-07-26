#!/usr/bin/env perl
use 5.020_000;
use warnings;
use Mojo::UserAgent qw<>;
use Mojo::JSON qw<decode_json>;
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

# Swapped to JSON at some point:
# const glyphs = { ... }
my $json = $page =~ m!<script>\s*const \s+ glyphs \s* = \s* ({.*?})\s*</script>!xms ? $1 : '';
die "Can't match const glyphs"
    if !length $json;

# "name" => "f0f0f0",
$json =~ s!,\s*}!}!xmsg;
my $href = decode_json($json);

open my $fh, '>', 'nf.go';
printf {$fh} "package main\n\n";
printf {$fh} "var nfRunes = map[rune]string{\n";
my %found;
for my $name (sort keys %$href) {
    my $hex = $href->{$name};
    $found{$hex}++ and next;
    if ($name =~ m!\Anfold-!xms) {
        printf {$fh} sprintf qq!\t%-8s"%s",\n!,
            (sprintf '%d:', hex("0x$hex")),
            ($name =~ s!\Anfold-!nf-!xmsr) . ' (obsolete)';
    } else {
        printf {$fh} sprintf qq!\t%-8s"%s",\n!,
            (sprintf '%d:', hex("0x$hex")),
            $name;
    }
}
printf {$fh} "}\n";
close $fh;
