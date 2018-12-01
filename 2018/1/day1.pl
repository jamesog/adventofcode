#!/usr/bin/env perl

use strict;
use warnings;

sub change_freq {
    my ($freq, $shift) = @_;
    $shift = int($shift);
    return $freq + $shift;
}

sub part_one {
    my @input = @_;
    my $freq = 0;
    foreach my $shift (@input) {
        $freq = change_freq($freq, $shift);
    }
    return $freq;
}

sub part_two {
    my @input = @_;
    my $freq = 0;
    my %seen_freqs;
    for (;;) {
        foreach my $shift (@input) {
            $freq = change_freq($freq, $shift);
            $seen_freqs{$freq}++;
            return $freq if ($seen_freqs{$freq} == 2);
        }
    }
}

my @input = <STDIN>;
print part_one(@input), "\n";
print part_two(@input), "\n";